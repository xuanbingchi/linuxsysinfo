package linuxsysinfo

// CPU information.
type CPUInfo struct {
	Cpu                string // cpu
	CpuModel           string // cpu model
	CpuVariation       string // cpu variation
	CpuRevision        string // cpu revision
	CpuSerialNumber    string // cpu serial number
	SystemType         string // system type
	SystemVariation    string // system variation
	SystemRevision     string // system revision
	SystemSerialNumber string // system serial number
	CPUFrequency       string // CPU frequency [MHz]
	PageSize           string // page size [bytes]
	PhysAddressBits    string // phys. address bits
	BogoMIPS           string // BogoMIPS
	KernelUnalignedAcc string // kernel unaligned acc
	UserUnalignedAcc   string // user unaligned acc
	PlatformString     string // platform string
	CpusDetected       string // cpus detected
	CpusActive         string // cpus active
	CpuActiveMask      string // cpu active mask
	CpusCore_start     string // cpus core_start
	L1Icache           string // L1 Icache
	L1Dcache           string // L1 Dcache
	L2Cache            string // L2 cache
	L3Cache            string // L3 cache
}

func CreatCPUInfo() (*CPUInfo, error) {
	i, err := getInfo1(ProcCpuinfoFile)
	if err != nil {
		return nil, err
	}

	c := new(CPUInfo)
	c.Cpu = i["cpu"]
	c.CpuModel = i["cpu model"]
	c.CpuVariation = i["cpu variation"]
	c.CpuRevision = i["cpu revision"]
	c.CpuSerialNumber = i["cpu serial number"]
	c.SystemType = i["system type"]
	c.SystemVariation = i["system variation"]
	c.SystemRevision = i["system revision"]
	c.SystemSerialNumber = i["system serial number"]
	c.CPUFrequency = i["CPU frequency [MHz]"]
	c.PageSize = i["page size [bytes]"]
	c.PhysAddressBits = i["phys. address bits"]
	c.BogoMIPS = i["BogoMIPS"]
	c.KernelUnalignedAcc = i["kernel unaligned acc"]
	c.UserUnalignedAcc = i["user unaligned acc"]
	c.PlatformString = i["platform string"]
	c.CpusDetected = i["cpus detected"]
	c.CpusActive = i["cpus active"]
	c.CpuActiveMask = i["cpu active mask"]
	c.CpusCore_start = i["cpus core_start"]
	c.L1Icache = i["L1 Icache"]
	c.L1Dcache = i["L1 Dcache"]
	c.L2Cache = i["L2 cache"]
	c.L3Cache = i["L3 cache"]

	return c, nil
}
