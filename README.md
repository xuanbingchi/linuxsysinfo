# linuxsysinfo

获取Linux系统信息

使用的系统文件：
```
/proc/meminfo
/proc/cpuinfo
```

使用的系统命令：
```
version
ifconfig
df -BKB -T
```

使用golang库：
```
net
```

##事例
###CPU信息
```
{
    "Cpu": "sw",
    "CpuModel": "sw",
    "CpuVariation": "1",
    "CpuRevision": "SW6A",
    "CpuSerialNumber": "",
    "SystemType": "Aere",
    "SystemVariation": "Aere",
    "SystemRevision": "0",
    "SystemSerialNumber": "",
    "CPUFrequency": "1600.00 ",
    "PageSize": "8192",
    "PhysAddressBits": "48",
    "BogoMIPS": "3200.00",
    "KernelUnalignedAcc": "1 (pc=ffffffff80e0a9b4,va=fffffc03b7005eba)",
    "UserUnalignedAcc": "1523214 (pc=20001033130,va=2000fc0eef3)",
    "PlatformString": "N/A",
    "CpusDetected": "16",
    "CpusActive": "16",
    "CpuActiveMask": "000000000000ffff",
    "CpusCore_start": "000000000000ffff",
    "L1Icache": "32K, 4-way, 128b line",
    "L1Dcache": "32K, 4-way, 128b line",
    "L2Cache": "512K, 8-way, 128b line",
    "L3Cache": "32768K, 8-way, 128b line"
}
```
###内存信息
```
{
    "MemTotal": 64883328,
    "MemFree": 53045344,
    "MemAvailable": 58279336,
    "Buffers": 323424,
    "Cached": 6138792,
    "SwapCached": 0,
    "Active": 10015376,
    "Inactive": 1169488,
    "Active_Anon": 4924520,
    "Inactive_Anon": 22936,
    "Active_File": 5090856,
    "Inactive_File": 1146552,
    "Unevictable": 0,
    "Mlocked": 0,
    "SwapTotal": 0,
    "SwapFree": 0,
    "Dirty": 1136,
    "Writeback": 0,
    "AnonPages": 4722776,
    "Mapped": 483152,
    "Shmem": 224816,
    "Slab": 361120,
    "SReclaimable": 245832,
    "SUnreclaim": 115288,
    "KernelStack": 19152,
    "PageTables": 34520,
    "NFS_Unstable": 0,
    "Bounce": 0,
    "WritebackTmp": 0,
    "CommitLimit": 32441664,
    "Committed_AS": 10055408,
    "VmallocTotal": 125829120,
    "VmallocUsed": 0,
    "VmallocChunk": 0,
    "AnonHugePages": 1916928,
    "HugePages_Total": 0,
    "HugePages_Free": 0,
    "HugePages_Rsvd": 0,
    "HugePages_Surp": 0,
    "Hugepagesize": 8192
  }
```
###系统版本信息
```
{
    "Name": "NeoKylin Server",
    "Platform": "5.0_U2",
    "Version": "sunway_64",
    "Release": "Release"
  }
```
###硬盘信息
```
[
    {
      "Filesystem": "devtmpfs",
      "Type": "devtmpfs",
      "Size": 33210377,
      "Used": 0,
      "Avail": 33210377,
      "Use": 0,
      "MountedOn": "/dev"
    },
    {
      "Filesystem": "tmpfs",
      "Type": "tmpfs   ",
      "Size": 33220264,
      "Used": 0,
      "Avail": 33220264,
      "Use": 0,
      "MountedOn": "/sys/fs/cgroup"
    },
    {
      "Filesystem": "/dev/sda1",
      "Type": "ext4    ",
      "Size": 31378572,
      "Used": 28781208,
      "Avail": 979833,
      "Use": 97,
      "MountedOn": "/"
    }
  ]
```
###IP信息
```
[
    {
      "Index": 1,
      "MTU": 65536,
      "Name": "lo",
      "HardwareAddr": "",
      "Flags": 5,
      "Ip": [
        {
          "Ip": "127.0.0.1",
          "SubnetMask": 8
        },
        {
          "Ip": "::1",
          "SubnetMask": 128
        }
      ]
    },
    {
      "Index": 2,
      "MTU": 1500,
      "Name": "enP1p38s0f0",
      "HardwareAddr": "6c:b3:11:41:65:90",
      "Flags": 19,
      "Ip": [
        {
          "Ip": "192.168.2.254",
          "SubnetMask": 24
        }
      ]
    },
    {
      "Index": 3,
      "MTU": 1500,
      "Name": "enP1p38s0f1",
      "HardwareAddr": "6c:b3:11:41:65:91",
      "Flags": 19,
      "Ip": [
        {
          "Ip": "192.168.3.254",
          "SubnetMask": 24
        }
      ]
    }
  ]
```