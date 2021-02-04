package linuxsysinfo

import (
	"bufio"
	"bytes"
	"regexp"
)

// CPU information.
type DfInfo struct {
	Filesystem string
	Type       string
	Size       int
	Used       int
	Avail      int
	Use        int
	MountedOn  string
}

func CreatDfInfo() ([]DfInfo, error) {
	txt, err := cmd("df -BKB -T")
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`^(\S+) +(.+) +([\d]+)kB +([\d]+)kB +([\d]+)kB +([\d]+)% +(\S+)$`)
	s := bufio.NewScanner(bytes.NewReader(txt))

	d := make([]DfInfo, 0, 1)
	isHead := true
	for s.Scan() {
		if isHead {
			isHead = false
			continue
		}
		p := r.FindStringSubmatch(s.Text())
		if p == nil || len(p) != 8 {
			continue
		}
		d = append(d, DfInfo{
			Filesystem: p[1],
			Type:       p[2],
			Size:       tryParseInt(p[3]),
			Used:       tryParseInt(p[4]),
			Avail:      tryParseInt(p[5]),
			Use:        tryParseInt(p[6]),
			MountedOn:  p[7],
		})
	}
	return d, nil
}
