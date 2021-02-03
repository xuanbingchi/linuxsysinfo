package linuxsysinfo

import (
	"bufio"
	"bytes"
	"os/exec"
)

type VersionInfo struct {
	Name     string
	Platform string
	Version  string
	Release  string
}

func CreatVersionInfo() (*VersionInfo, error) {
	v := new(VersionInfo)
	command := exec.Command("/bin/bash", "-c", "version")
	txt, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(bytes.NewReader(txt))
	index := 0
	for s.Scan() {
		t := s.Text()
		if len(t) != 0 && t[0] == '-' {
			switch index {
			case 0:
				v.Name = t[1:]
				index++
			case 1:
				v.Platform = t[1:]
				index++
			case 2:
				v.Version = t[1:]
				index++
			case 3:
				v.Release = t[1:]
				index++
			}
		}

	}

	return v, nil
}
