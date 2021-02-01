package linuxsysinfo

import (
	"bufio"
	"bytes"
	"os/exec"
	"regexp"
)

var re = []*regexp.Regexp{
	0:  regexp.MustCompile(`^([\w-]+): flags=[\d]+<([\w,]+)> +mtu ([\d]+)$`),
	1:  regexp.MustCompile(`^ +inet ([\d.]+) +netmask ([\d.]+) +broadcast ([\d.]+)$`),
	2:  regexp.MustCompile(`^ +inet ([\d.]+) +netmask ([\d.]+)$`),
	3:  regexp.MustCompile(`^ +inet6 ([\d\w:]+) +prefixlen ([\d]+) +scopeid ([\d\w.<>]+)$`),
	4:  regexp.MustCompile(`^ +ether ([\d\w:]+) +txqueuelen ([\d]+) +\(([\w]+)\)$`),
	5:  regexp.MustCompile(`^ +loop (.+)$`),
	6:  regexp.MustCompile(`^ +RX packets ([\d]+) +bytes ([\d]+) \((.+)\)$`),
	7:  regexp.MustCompile(`^ +RX errors ([\d]+) +dropped ([\d]+) +overruns ([\d]+) +frame ([\d]+)$`),
	8:  regexp.MustCompile(`^ +TX packets ([\d]+) +bytes ([\d]+) \((.+)\)$`),
	9:  regexp.MustCompile(`^ +TX errors +([\d]+) +dropped +([\d]+) +overruns +([\d]+) +carrier +([\d]+) +collisions ([\d]+)$`),
	10: regexp.MustCompile(`^ +device memory ([\d\w-]+)$`),
}

type IfConfigInfo struct {
	Name  string
	Flags string
	Mtu   string

	Inet      string
	Netmask   string
	Broadcast string

	Inet6     string
	Prefixlen string
	Scopeid   string

	Ether          string
	Txqueuelen     string
	Txqueuelen_des string
	Loop           string

	RX_Packets       string
	RX_Packets_bytes string
	RX_Packets_des   string
	RX_errors        string
	RX_dropped       string
	RX_overruns      string
	RX_frame         string

	TX_Packets       string
	TX_Packets_bytes string
	TX_Packets_des   string
	TX_errors        string
	TX_dropped       string
	TX_overruns      string
	TX_carrier       string
	TX_collisions    string

	Device_memory string
}

func CreatIfConfigInfos() ([]IfConfigInfo, error) {
	command := exec.Command("/bin/bash", "-c", "ifconfig")
	txt, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}

	tt := bytes.Split(txt, []byte("\n\n"))
	i := make([]IfConfigInfo, 0, len(tt))
	for _, s := range tt {
		i = append(i, *ifConfigParse(s))
	}

	return i, nil
}

func ifConfigParse(txt []byte) *IfConfigInfo {
	i := new(IfConfigInfo)

	s := bufio.NewScanner(bytes.NewReader(txt))
	var isContinue = false
	var t string
	for index, r := range re {
		if !isContinue {
			if !s.Scan() {
				return i
			}
			t = s.Text()
		}
		p := r.FindStringSubmatch(t)
		isContinue = p == nil
		if isContinue {
			continue
		}

		switch index {
		case 0:
			i.Name = p[1]
			i.Flags = p[2]
			i.Mtu = p[3]
		case 1:
			i.Inet = p[1]
			i.Netmask = p[2]
			i.Broadcast = p[3]
		case 2:
			i.Inet = p[1]
			i.Netmask = p[2]
		case 3:
			i.Inet6 = p[1]
			i.Prefixlen = p[2]
			i.Scopeid = p[3]
		case 4:
			i.Ether = p[1]
			i.Txqueuelen = p[2]
			i.Txqueuelen_des = p[3]
		case 5:
			i.Loop = p[1]
		case 6:
			i.RX_Packets = p[1]
			i.RX_Packets_bytes = p[2]
			i.RX_Packets_des = p[3]
		case 7:
			i.RX_errors = p[1]
			i.RX_dropped = p[2]
			i.RX_overruns = p[3]
			i.RX_frame = p[4]
		case 8:
			i.TX_Packets = p[1]
			i.TX_Packets_bytes = p[2]
			i.TX_Packets_des = p[3]
		case 9:
			i.TX_errors = p[1]
			i.TX_dropped = p[2]
			i.TX_overruns = p[3]
			i.TX_carrier = p[4]
			i.TX_collisions = p[5]
		case 10:
			i.Device_memory = p[1]
		}
	}

	return i
}
