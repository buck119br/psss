package psss

import (
	"fmt"
	"math"
	"os"
	"syscall"
)

type FileInfo struct {
	Path    string
	Name    string
	SysStat *syscall.Stat_t
}

func NewFileInfo() *FileInfo {
	fi := new(FileInfo)
	return fi
}

func (fi *FileInfo) GetStat(path string, name string) (err error) {
	stat, err := os.Stat(path + "/" + name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fi.Path = path
	fi.Name = name
	var ok bool
	if fi.SysStat, ok = stat.Sys().(*syscall.Stat_t); !ok {
		fmt.Printf("FileInfo.Sys:[%v] assertion failure\n", stat)
		return fmt.Errorf("FileInfo.Sys:[%v] assertion failure", stat)
	}
	return nil
}

func BwToStr(bw float64) string {
	switch {
	case bw > math.Pow(1000, 7):
		return fmt.Sprintf("%.2fZ", bw/math.Pow(1000, 7))
	case bw > math.Pow(1000, 6):
		return fmt.Sprintf("%.2fE", bw/math.Pow(1000, 6))
	case bw > math.Pow(1000, 5):
		return fmt.Sprintf("%.2fP", bw/math.Pow(1000, 5))
	case bw > math.Pow(1000, 4):
		return fmt.Sprintf("%.2fT", bw/math.Pow(1000, 4))
	case bw > math.Pow(1000, 3):
		return fmt.Sprintf("%.2fG", bw/math.Pow(1000, 3))
	case bw > math.Pow(1000, 2):
		return fmt.Sprintf("%.2fM", bw/math.Pow(1000, 2))
	case bw > math.Pow(1000, 1):
		return fmt.Sprintf("%.2fK", bw/math.Pow(1000, 1))
	}
	return fmt.Sprintf("%g", bw)
}
