package linuxsysinfo

// Mem information.
type MemInfo struct {
	MemTotal        string //MemTotal
	MemFree         string //MemFree
	MemAvailable    string //MemAvailable
	Buffers         string //Buffers
	Cached          string //Cached
	SwapCached      string //SwapCached
	Active          string //Active
	Inactive        string //Inactive
	Active_Anon     string //Active(anon)
	Inactive_Anon   string //Inactive(anon)
	Active_File     string //Active(file)
	Inactive_File   string //Inactive(file)
	Unevictable     string //Unevictable
	Mlocked         string //Mlocked
	SwapTotal       string //SwapTotal
	SwapFree        string //SwapFree
	Dirty           string //Dirty
	Writeback       string //Writeback
	AnonPages       string //AnonPages
	Mapped          string //Mapped
	Shmem           string //Shmem
	Slab            string //Slab
	SReclaimable    string //SReclaimable
	SUnreclaim      string //SUnreclaim
	KernelStack     string //KernelStack
	PageTables      string //PageTables
	NFS_Unstable    string //NFS_Unstable
	Bounce          string //Bounce
	WritebackTmp    string //WritebackTmp
	CommitLimit     string //CommitLimit
	Committed_AS    string //Committed_AS
	VmallocTotal    string //VmallocTotal
	VmallocUsed     string //VmallocUsed
	VmallocChunk    string //VmallocChunk
	AnonHugePages   string //AnonHugePages
	HugePages_Total string //HugePages_Total
	HugePages_Free  string //HugePages_Free
	HugePages_Rsvd  string //HugePages_Rsvd
	HugePages_Surp  string //HugePages_Surp
	Hugepagesize    string //Hugepagesize
}

func CreatMemInfo() (*MemInfo, error) {
	i, err := getInfo2(ProcMeminfoFile)
	if err != nil {
		return nil, err
	}

	c := new(MemInfo)
	c.MemTotal = i["MemTotal"]
	c.MemFree = i["MemFree"]
	c.MemAvailable = i["MemAvailable"]
	c.Buffers = i["Buffers"]
	c.Cached = i["Cached"]
	c.SwapCached = i["SwapCached"]
	c.Active = i["Active"]
	c.Inactive = i["Inactive"]
	c.Active_Anon = i["Active(anon)"]
	c.Inactive_Anon = i["Inactive(anon)"]
	c.Active_File = i["Active(file)"]
	c.Inactive_File = i["Inactive(file)"]
	c.Unevictable = i["Unevictable"]
	c.Mlocked = i["Mlocked"]
	c.SwapTotal = i["SwapTotal"]
	c.SwapFree = i["SwapFree"]
	c.Dirty = i["Dirty"]
	c.Writeback = i["Writeback"]
	c.AnonPages = i["AnonPages"]
	c.Mapped = i["Mapped"]
	c.Shmem = i["Shmem"]
	c.Slab = i["Slab"]
	c.SReclaimable = i["SReclaimable"]
	c.SUnreclaim = i["SUnreclaim"]
	c.KernelStack = i["KernelStack"]
	c.PageTables = i["PageTables"]
	c.NFS_Unstable = i["NFS_Unstable"]
	c.Bounce = i["Bounce"]
	c.WritebackTmp = i["WritebackTmp"]
	c.CommitLimit = i["CommitLimit"]
	c.Committed_AS = i["Committed_AS"]
	c.VmallocTotal = i["VmallocTotal"]
	c.VmallocUsed = i["VmallocUsed"]
	c.VmallocChunk = i["VmallocChunk"]
	c.AnonHugePages = i["AnonHugePages"]
	c.HugePages_Total = i["HugePages_Total"]
	c.HugePages_Free = i["HugePages_Free"]
	c.HugePages_Rsvd = i["HugePages_Rsvd"]
	c.HugePages_Surp = i["HugePages_Surp"]
	c.Hugepagesize = i["Hugepagesize"]

	return c, nil
}
