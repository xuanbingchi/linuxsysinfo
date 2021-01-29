package linuxsysinfo

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func ifconfigParse(txt string) (*IfconfigInfo, error) {

	i := new(IfconfigInfo)

	s := bufio.NewScanner(strings.NewReader(txt))
	index := 0
	for s.Scan() {
		t := s.Text()
		switch index {
		case 0:
			m := regexp.MustCompile(`^([\w-]+): flags=[\d]+<([\w,]+)>  mtu ([\d]+)$`)
			p := m.FindStringSubmatch(t)
			if len(p) != 4 {
				return nil, fmt.Errorf("parse ifconfig failed.line:%d param number:%d", index, len(p))
			}
			i.Name = p[1]
			i.Flags = p[2]
			i.Mtu = p[3]
		case 1:
			m := regexp.MustCompile(`^ +inet ([\d.]+)  netmask ([\d.]+)  broadcast ([\d.]+)$`)
			p := m.FindStringSubmatch(t)
			if len(p) != 4 {
				return nil, fmt.Errorf("parse ifconfig failed.line:%d param number:%d", index, len(p))
			}
			i.Inet = p[1]
			i.Netmask = p[2]
			i.Broadcast = p[3]
		case 2:

		case 3:
		case 4:
		case 5:
		case 6:
		case 7:

		}
		index++
	}

	return nil, nil
}

type IfconfigInfo struct {
	Name      string
	Flags     string
	Mtu       string
	Inet      string
	Netmask   string
	Broadcast string
}
