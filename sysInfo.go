package linuxsysinfo

import (
	"bufio"
	"os"
	"regexp"
)

const (
	ProcMeminfoFile = "/proc/meminfo"
	ProcCpuinfoFile = "/proc/cpuinfo"
)

type info map[string]string

func getInfo0(file string, fu func(text string)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fu(s.Text())
	}
	return nil
}

func getInfo1(file string) (info, error) {
	info := make(info)
	m := regexp.MustCompile("\t+: ")

	err := getInfo0(file, func(text string) {
		sl := m.Split(text, 2)
		if sl != nil && len(sl) == 2 {
			info[sl[0]] = sl[1]
		}
	})

	return info, err
}

func getInfo2(file string) (info, error) {
	info := make(info)
	m := regexp.MustCompile(`^([\w_\(\)]+): +([\d]+)[\D]*$`)
	err := getInfo0(file, func(text string) {
		sl := m.FindStringSubmatch(text)
		if sl != nil && len(sl) == 3 {
			info[sl[1]] = sl[2]
		}
	})

	return info, err
}
