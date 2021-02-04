package linuxsysinfo

// Mem information.
type MemInfo struct {
	MemTotal        int //MemTotal
	MemFree         int //MemFree
	MemAvailable    int //MemAvailable
	Buffers         int //Buffers
	Cached          int //Cached
	SwapCached      int //SwapCached
	Active          int //Active
	Inactive        int //Inactive
	Active_Anon     int //Active(anon)
	Inactive_Anon   int //Inactive(anon)
	Active_File     int //Active(file)
	Inactive_File   int //Inactive(file)
	Unevictable     int //Unevictable
	Mlocked         int //Mlocked
	SwapTotal       int //SwapTotal
	SwapFree        int //SwapFree
	Dirty           int //Dirty
	Writeback       int //Writeback
	AnonPages       int //AnonPages
	Mapped          int //Mapped
	Shmem           int //Shmem
	Slab            int //Slab
	SReclaimable    int //SReclaimable
	SUnreclaim      int //SUnreclaim
	KernelStack     int //KernelStack
	PageTables      int //PageTables
	NFS_Unstable    int //NFS_Unstable
	Bounce          int //Bounce
	WritebackTmp    int //WritebackTmp
	CommitLimit     int //CommitLimit
	Committed_AS    int //Committed_AS
	VmallocTotal    int //VmallocTotal
	VmallocUsed     int //VmallocUsed
	VmallocChunk    int //VmallocChunk
	AnonHugePages   int //AnonHugePages
	HugePages_Total int //HugePages_Total
	HugePages_Free  int //HugePages_Free
	HugePages_Rsvd  int //HugePages_Rsvd
	HugePages_Surp  int //HugePages_Surp
	Hugepagesize    int //Hugepagesize
}

func CreatMemInfo() (*MemInfo, error) {
	i, err := getInfo2(ProcMeminfoFile)
	if err != nil {
		return nil, err
	}

	c := new(MemInfo)
	c.MemTotal = tryParseInt(i["MemTotal"])
	c.MemFree = tryParseInt(i["MemFree"])
	c.MemAvailable = tryParseInt(i["MemAvailable"])
	c.Buffers = tryParseInt(i["Buffers"])
	c.Cached = tryParseInt(i["Cached"])
	c.SwapCached = tryParseInt(i["SwapCached"])
	c.Active = tryParseInt(i["Active"])
	c.Inactive = tryParseInt(i["Inactive"])
	c.Active_Anon = tryParseInt(i["Active(anon)"])
	c.Inactive_Anon = tryParseInt(i["Inactive(anon)"])
	c.Active_File = tryParseInt(i["Active(file)"])
	c.Inactive_File = tryParseInt(i["Inactive(file)"])
	c.Unevictable = tryParseInt(i["Unevictable"])
	c.Mlocked = tryParseInt(i["Mlocked"])
	c.SwapTotal = tryParseInt(i["SwapTotal"])
	c.SwapFree = tryParseInt(i["SwapFree"])
	c.Dirty = tryParseInt(i["Dirty"])
	c.Writeback = tryParseInt(i["Writeback"])
	c.AnonPages = tryParseInt(i["AnonPages"])
	c.Mapped = tryParseInt(i["Mapped"])
	c.Shmem = tryParseInt(i["Shmem"])
	c.Slab = tryParseInt(i["Slab"])
	c.SReclaimable = tryParseInt(i["SReclaimable"])
	c.SUnreclaim = tryParseInt(i["SUnreclaim"])
	c.KernelStack = tryParseInt(i["KernelStack"])
	c.PageTables = tryParseInt(i["PageTables"])
	c.NFS_Unstable = tryParseInt(i["NFS_Unstable"])
	c.Bounce = tryParseInt(i["Bounce"])
	c.WritebackTmp = tryParseInt(i["WritebackTmp"])
	c.CommitLimit = tryParseInt(i["CommitLimit"])
	c.Committed_AS = tryParseInt(i["Committed_AS"])
	c.VmallocTotal = tryParseInt(i["VmallocTotal"])
	c.VmallocUsed = tryParseInt(i["VmallocUsed"])
	c.VmallocChunk = tryParseInt(i["VmallocChunk"])
	c.AnonHugePages = tryParseInt(i["AnonHugePages"])
	c.HugePages_Total = tryParseInt(i["HugePages_Total"])
	c.HugePages_Free = tryParseInt(i["HugePages_Free"])
	c.HugePages_Rsvd = tryParseInt(i["HugePages_Rsvd"])
	c.HugePages_Surp = tryParseInt(i["HugePages_Surp"])
	c.Hugepagesize = tryParseInt(i["Hugepagesize"])

	return c, nil
}
