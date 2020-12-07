package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	fmt.Println(diskInfo.Total / GB)
	fmt.Println(diskInfo.Free / GB)
	fmt.Println(diskInfo.UsedPercent)
}
