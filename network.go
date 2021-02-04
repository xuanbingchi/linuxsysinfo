package linuxsysinfo

import (
	"net"
	"strings"
)

type NetInfo struct {
	Index        int
	MTU          int
	Name         string
	HardwareAddr string
	Flags        net.Flags
	Ip           []Ip
}

type Ip struct {
	Ip         string
	SubnetMask int
}

func CreatNetInfo() ([]NetInfo, error) {
	a, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	n := make([]NetInfo, 0, len(a))

	for _, i := range a {
		nn := NetInfo{
			Index:        i.Index,
			MTU:          i.MTU,
			Name:         i.Name,
			HardwareAddr: i.HardwareAddr.String(),
			Flags:        i.Flags,
		}

		add, err := i.Addrs()
		if err != nil {
			continue
		}
		nn.Ip = make([]Ip, 0, len(add))
		for _, addr := range add {
			s := strings.Split(addr.String(), "/")
			var ip Ip
			if len(s) == 2 {
				ip.Ip = s[0]
				ip.SubnetMask = tryParseInt(s[1])
			} else if len(s) == 1 {
				ip.Ip = s[0]
			}
			nn.Ip = append(nn.Ip, ip)
		}

		n = append(n, nn)
	}
	return n, nil
}
