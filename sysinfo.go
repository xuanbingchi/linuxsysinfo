package linuxsysinfo

import (
	"bufio"
	"os"
	"regexp"
)

const (
	//ProcMeminfoFile     = "/proc/meminfo"
	//ProcCpuinfoFile     = "/proc/cpuinfo"
	//ProcLoadavgFile     = "/proc/loadavg"
	//ProcFilesystemsFile = "/proc/filesystems"

	ProcMeminfoFile     = "/proc/meminfo"
	ProcCpuinfoFile     = `C:\gopath\work\linuxsysinfo\example\cpuinfo`
	ProcLoadavgFile     = "/proc/loadavg"
	ProcFilesystemsFile = "/proc/filesystems"
)

var (
	reTwoColumns = regexp.MustCompile("\t+: ")
	reExtraSpace = regexp.MustCompile(" +")
	reCacheSize  = regexp.MustCompile(`^(\d+) KB$`)
)

type info map[string]string

func getTwoColumns(file string) (info, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	info := make(info)
	s := bufio.NewScanner(f)
	for s.Scan() {
		sl := reTwoColumns.Split(s.Text(), 2)
		if sl != nil && len(sl) == 2 {
			info[sl[0]] = sl[1]
		}
	}
	return info, nil
}
