package probe

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var GConfig *ProbeConfig

type ProbeConfig struct {
	SamplingFrequency uint64
	TransmitInterval  uint64

	Items struct {
		IO struct {
			NIC struct {
				Switch     bool
				Interfaces []string
			}
		}
		FileSystem struct {
			MountInfo struct {
				Switch        bool
				MountPoints   []string
				MountPointSet map[string]bool
			}
			FileInfo struct {
				Switch   bool
				FilePath []string
			}
		}
		Process struct {
			Switch      bool
			ProcName    []string
			ProcNameSet map[string]bool
		}
		Watchdog struct {
			Switch bool
			Ports  []struct {
				Name     string
				Protocol string
				Addr     string
				Port     string
			}
		}
		Camera struct {
			Switch bool
		}
	}
}

func (pc *ProbeConfig) Load(path string) error {
	if _, err := toml.DecodeFile(path, pc); err != nil {
		return err
	}
	if pc.Items.FileSystem.MountInfo.Switch && len(pc.Items.FileSystem.MountInfo.MountPoints) > 0 {
		pc.Items.FileSystem.MountInfo.MountPointSet = make(map[string]bool)
		for _, v := range pc.Items.FileSystem.MountInfo.MountPoints {
			pc.Items.FileSystem.MountInfo.MountPointSet[v] = true
		}
	}
	if pc.Items.Process.Switch && len(pc.Items.Process.ProcName) > 0 {
		pc.Items.Process.ProcNameSet = make(map[string]bool)
		for _, v := range pc.Items.Process.ProcName {
			pc.Items.Process.ProcNameSet[v] = true
		}
	}
	return nil
}

func (pc *ProbeConfig) Check() error {
	if pc.TransmitInterval == 0 {
		return fmt.Errorf("invalid transmit interval")
	}
	if pc.SamplingFrequency == 0 {
		return fmt.Errorf("invalid sampling frequency")
	}
	if pc.TransmitInterval%pc.SamplingFrequency != 0 {
		return fmt.Errorf("transmit interval should be divisible by sampling frequency")
	}
	if pc.TransmitInterval/pc.SamplingFrequency < 5 {
		return fmt.Errorf("too fast sampling")
	}
	return nil
}
