package probe

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

var GProbe Probe = newProbe()

type Probe interface {
	Init(config *ProbeConfig) error
	GetContext() (*ProbeContext, error)
}

type probe struct {
	mutex sync.Mutex

	kChan chan int

	ctx *ProbeContext
}

func newProbe() *probe {
	p := new(probe)
	p.kChan = make(chan int)
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
	for sig := range p.kChan {
		switch sig {
		case 1:
			go p.samplingTimer()
		}
	}
}

func (p *probe) samplingTimer() {
	interval := time.Duration(GConfig.SamplingInterval) * time.Second
	now := time.Now()
	timer := time.NewTimer(now.Truncate(interval).Add(interval).Sub(time.Now()))

	defer func() {
		if rcvErr := recover(); rcvErr != nil {
			logger.Errorf("recovered from panic error:[%v]", rcvErr)
			debug.PrintStack()
		}
		timer.Stop()
		p.kChan <- 1
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
			p.ctx.Fit(newCtx)
			p.mutex.Unlock()

			logger.WithField("time_cost", time.Since(now).Seconds()).Infof("finished")
		}
	}
}

func (p *probe) Init(config *ProbeConfig) error {
	GConfig = config
	err := GConfig.Check()
	if err != nil {
		return err
	}

	p.ctx = NewProbeContext()

	go p.keeper()
	go p.samplingTimer()

	logger.WithField("config", GConfig).Info("finished")

	return nil
}

func (p *probe) GetContext() (*ProbeContext, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ctx := p.ctx
	p.ctx = NewProbeContext()

	if ctx.SamplingCounter < 1 {
		return nil, fmt.Errorf("too few sample")
	}

	ctx.Average()

	return ctx, nil
}
