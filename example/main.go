package main

import (
	"encoding/xml"
	"fmt"
	"linuxsysinfo"
)

func main() {
	type info struct {
		CPUInfo       *linuxsysinfo.CPUInfo
		MemInfo       *linuxsysinfo.MemInfo
		IfConfigInfos []linuxsysinfo.IfConfigInfo
	}

	i := new(info)
	var err error
	i.CPUInfo, err = linuxsysinfo.CreatCPUInfo()
	if err != nil {
		panic(err)
	}
	i.MemInfo, err = linuxsysinfo.CreatMemInfo()
	if err != nil {
		panic(err)
	}
	i.IfConfigInfos, err = linuxsysinfo.CreatIfConfigInfos()
	if err != nil {
		panic(err)
	}

	ii, err := xml.Marshal(i)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(ii))

}
