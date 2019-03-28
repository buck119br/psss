package probe

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var GConfig *ProbeConfig

type ProbeConfig struct {
	SamplingInterval uint64

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
}

func (pc *ProbeConfig) Load(path string) error {
	if _, err := toml.DecodeFile(path, pc); err != nil {
		return err
	}
	if pc.FileSystem.MountInfo.Switch && len(pc.FileSystem.MountInfo.MountPoints) > 0 {
		pc.FileSystem.MountInfo.MountPointSet = make(map[string]bool)
		for _, v := range pc.FileSystem.MountInfo.MountPoints {
			pc.FileSystem.MountInfo.MountPointSet[v] = true
		}
	}
	if pc.Process.Switch && len(pc.Process.ProcName) > 0 {
		pc.Process.ProcNameSet = make(map[string]bool)
		for _, v := range pc.Process.ProcName {
			pc.Process.ProcNameSet[v] = true
		}
	}
	return nil
}

func (pc *ProbeConfig) Check() error {
	if pc.SamplingInterval < 5 {
		return fmt.Errorf("too fast sampling")
	}
	return nil
}
