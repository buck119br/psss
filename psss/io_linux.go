// +build linux

package psss

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// definition comes from Linux kernel /Documentation/iostats.txt
type DiskStat struct {
	MajorNumber      uint64
	MinorNumber      uint64
	Name             string
	ReadCompleted    uint64 // reads completed
	ReadMerged       uint64 // reads merged, field 6 -- # of writes merged
	SectorsRead      uint64 // sectors read
	ReadingSpent     uint64 // milliseconds spent reading
	WriteCompleted   uint64 // writes completed
	WriteMerged      uint64 // writes merged
	SectorsWritten   uint64 // sectors written
	WritingSpent     uint64 // milliseconds spent writing
	IOProgressing    uint64 // I/Os currently in progress
	IOSpent          uint64 // milliseconds spent doing I/Os
	WeightedIOSpent  uint64 // milliseconds spent doing I/Os (weighted)
	DiscardCompleted uint64 // discards completed
	DiscardMerged    uint64 // discards merged
	SectorDiscarded  uint64 // sectors discarded
	DiscardSpending  uint64 // milliseconds spent discarding
}

func (ds *DiskStat) Parse(numFields int, raw string) (err error) {
	fields := strings.Fields(SlimSpaceRegExp.ReplaceAllString(raw, " "))

	if ds.MajorNumber, err = strconv.ParseUint(fields[0], 10, 64); err != nil {
		return err
	}
	if ds.MinorNumber, err = strconv.ParseUint(fields[1], 10, 64); err != nil {
		return err
	}
	ds.Name = fields[2]

	fields = fields[3:]

	var v uint64
	for i := 0; i < numFields; i++ {
		if v, err = strconv.ParseUint(fields[i], 10, 64); err != nil {
			return err
		}
		switch i {
		case 0:
			ds.ReadCompleted = v
		case 1:
			ds.ReadMerged = v
		case 2:
			ds.SectorsRead = v
		case 3:
			ds.ReadingSpent = v
		case 4:
			ds.WriteCompleted = v
		case 5:
			ds.WriteMerged = v
		case 6:
			ds.SectorsWritten = v
		case 7:
			ds.WritingSpent = v
		case 8:
			ds.IOProgressing = v
		case 9:
			ds.IOSpent = v
		case 10:
			ds.WeightedIOSpent = v
		case 11:
			ds.DiscardCompleted = v
		case 12:
			ds.DiscardMerged = v
		case 13:
			ds.SectorDiscarded = v
		case 14:
			ds.DiscardSpending = v
		default:
			return fmt.Errorf("unknown field index:[%d]", i)
		}
	}

	return nil
}

type DiskStats []*DiskStat

func NewDiskStats() DiskStats {
	return make([]*DiskStat, 0)
}

func (dss *DiskStats) Get() (err error) {
	var numFields int

	switch {
	case KVer.VersionNum < 2006000:
		return fmt.Errorf("kernel version:[%s] not support", KVer.UTSRelease)
	case KVer.VersionNum >= 2006000 && KVer.VersionNum < 4018000:
		numFields = 11
	case KVer.VersionNum >= 4018000:
		numFields = 15
	}

	fd, err := os.Open(ProcRoot + "/diskstats")
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		ds := new(DiskStat)
		if err = ds.Parse(numFields, scanner.Text()); err != nil {
			fmt.Printf("disk stat parse error:[%v]\n", err)
			continue
		}
		*dss = append(*dss, ds)
	}

	return nil
}

type NetDev struct {
	Interface          string
	ReceiveBytes       uint64
	ReceivePackets     uint64
	ReceiveErrs        uint64
	ReceiveDrop        uint64
	ReceiveFifo        uint64
	ReceiveFrame       uint64
	ReceiveCompressed  uint64
	ReceiveMulticast   uint64
	TransmitBytes      uint64
	TransmitPackets    uint64
	TransmitErrs       uint64
	TransmitDrop       uint64
	TransmitFifo       uint64
	TransmitColls      uint64
	TransmitCarrier    uint64
	TransmitCompressed uint64
}

func (nd *NetDev) Parse(raw string) (err error) {
	fields := strings.Split(SlimSpaceRegExp.ReplaceAllString(raw, " "), ":")
	nd.Interface = strings.TrimSpace(fields[0])

	var fCtr int
	var v uint64

	fCtr = 0
	for _, s := range strings.Fields(fields[1]) {
		if len(s) == 0 {
			continue
		}
		if v, err = strconv.ParseUint(s, 10, 64); err != nil {
			return fmt.Errorf("parse field:[%s] error:[%v]", s, err)
		}
		switch fCtr {
		case 0:
			nd.ReceiveBytes = v
		case 1:
			nd.ReceivePackets = v
		case 2:
			nd.ReceiveErrs = v
		case 3:
			nd.ReceiveDrop = v
		case 4:
			nd.ReceiveFifo = v
		case 5:
			nd.ReceiveFrame = v
		case 6:
			nd.ReceiveCompressed = v
		case 7:
			nd.ReceiveMulticast = v
		case 8:
			nd.TransmitBytes = v
		case 9:
			nd.TransmitPackets = v
		case 10:
			nd.TransmitErrs = v
		case 11:
			nd.TransmitDrop = v
		case 12:
			nd.TransmitFifo = v
		case 13:
			nd.TransmitColls = v
		case 14:
			nd.TransmitCarrier = v
		case 15:
			nd.TransmitCompressed = v
		default:
			return fmt.Errorf("invalid field:[%s]", s)
		}
		fCtr++
	}
	return nil
}

type NetDevs map[string]*NetDev

func NewNetDevs() NetDevs {
	return make(map[string]*NetDev)
}

func (nds *NetDevs) Get() error {
	fd, err := os.Open(ProcRoot + "/self/net/dev")
	if err != nil {
		return err
	}
	defer fd.Close()

	var lCtr int
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return err
		}
		lCtr++
		if lCtr < 3 {
			continue
		}

		nd := new(NetDev)
		if err = nd.Parse(scanner.Text()); err != nil {
			fmt.Printf("net dev parse error:[%v]\n", err)
			continue
		}
		if nd.Interface == "lo" {
			continue
		}
		(*nds)[nd.Interface] = nd
	}
	return nil
}
