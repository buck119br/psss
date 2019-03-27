package probe

import (
	"runtime/debug"
	"sync"
	"time"
)

var GProbe Probe = newProbe()

type Probe interface {
	Init(configPath string) error
}

type probe struct {
	mutex sync.Mutex

	keeperChan chan int

	ctxFitted *ProbeContext
}

func newProbe() *probe {
	p := new(probe)
	p.keeperChan = make(chan int)
	return p
}

func (p *probe) keeper() {
	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic error:[%v]", rcvErr)
			debug.PrintStack()
		}
	}()
	logger.Info("started")
	for sig := range p.keeperChan {
		switch sig {
		case 1:
			go p.samplingTimer()
		case 2:
			go p.transmitTimer()
		}
	}
}

func (p *probe) samplingTimer() {
	interval := time.Duration(GConfig.TransmitInterval/GConfig.SamplingFrequency) * time.Second
	now := time.Now()
	timer := time.NewTimer(now.Truncate(interval).Add(interval).Sub(time.Now()))

	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic error:[%v]", rcvErr)
			debug.PrintStack()
		}
		timer.Stop()
		p.keeperChan <- 1
	}()
	logger.Info("started")

	var err error
	for {
		timer.Reset(now.Truncate(interval).Add(interval).Sub(time.Now()))
		select {
		case now = <-timer.C:
			now = now.Truncate(time.Second)

			newCtx := NewProbeContext()
			if err = newCtx.Sample(); err != nil {
				logger.Errorf("sample error:[%v]", err)
				continue
			}

			p.mutex.Lock()
			p.ctxFitted.Fit(newCtx)
			logger.WithField("ctx", p.ctxFitted).Infof("haha")
			p.mutex.Unlock()

			logger.WithField("time_cost", time.Since(now).Seconds()).Infof("finished")
		}
	}
}

func (p *probe) transmitTimer() {
	interval := time.Duration(GConfig.TransmitInterval) * time.Second
	now := time.Now()
	timer := time.NewTimer(now.Truncate(interval).Add(interval).Sub(time.Now()))

	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic error:[%v]", rcvErr)
			debug.PrintStack()
		}
		p.keeperChan <- 2
	}()
	logger.Info("started")

	var tmpCtx *ProbeContext
	for {
		timer.Reset(now.Truncate(interval).Add(interval).Sub(time.Now()))
		select {
		case now = <-timer.C:
			now = now.Truncate(time.Second)

			p.mutex.Lock()
			tmpCtx = p.ctxFitted
			p.ctxFitted = NewProbeContext()
			p.mutex.Unlock()

			if tmpCtx.samplingCounter == 0 {
				logger.Warnf("too few samples")
				continue
			}
		}
	}
}

func (p *probe) Init(configPath string) error {
	GConfig = new(ProbeConfig)
	err := GConfig.Load(configPath)
	if err != nil {
		return err
	}
	if err = GConfig.Check(); err != nil {
		return err
	}

	p.ctxFitted = NewProbeContext()

	go p.keeper()
	go p.samplingTimer()
	go p.transmitTimer()

	logger.WithField("config", GConfig).Info("finished")

	return nil
}
