package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	GB  = 1000 * 1000 * 1000
	GiB = 1024 * 1024 * 1024
)

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

func PrintMem() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v GiB, Free: %v MiB, UsedPercent: %f%%\n", v.Total/GiB, v.Free/1024/1024, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println()
	fmt.Println(v)
	fmt.Println()
}

func PrintDisk() {
	info3, _ := disk.IOCounters() //所有硬盘的io信息
	fmt.Println("info:", info3)

	parts, _ := disk.Partitions(false)

	for _, p := range parts {
		diskInfo, _ := disk.Usage(p.Mountpoint)

		size := float64(diskInfo.Total)
		fmt.Println("path:", diskInfo.Path, p)
		fmt.Println("all:", p.Device, p.Fstype, p.Mountpoint, size/GiB, size/GB)
		fmt.Println("used:", float64(diskInfo.Used)/GiB, float64(diskInfo.Used)/GB)
		fmt.Println("free:", float64(diskInfo.Free)/GiB, float64(diskInfo.Free)/GB)

		//fmt.Println(1, p.String())
		//fmt.Println(2, diskInfo.String())
	}
}

func main() {
	fmt.Println("cpu: ", GetCpuPercent())
	fmt.Println("mem: ", GetMemPercent())
	fmt.Println("disk:", GetDiskPercent())

	fmt.Println()
	PrintMem()
	PrintDisk()
}
