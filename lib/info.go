package lib

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func CpuInfo() {
	cpu, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting host information:", err)
		return
	}
	fmt.Printf("%sVendor ID:%s \t\t%s\n", ColorGreen, ColorNone, cpu[0].VendorID)
	var cores int
	for _, s := range cpu {
		cores = int(s.Cores) + cores
	}
	fmt.Printf("%sCores:%s \t\t\t%v\n", ColorGreen, ColorNone, cores)
	fmt.Printf("%sMhz:%s \t\t\t%v\n", ColorGreen, ColorNone, cpu[0].Mhz)
	fmt.Printf("%sModel Name:%s \t\t%v\n", ColorGreen, ColorNone, cpu[0].ModelName)
}

func HostInfo() {
	host, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host information:", err)
		return
	}
	fmt.Printf("%sHostname:%s \t\t%s\n", ColorGreen, ColorNone, host.Hostname)
	fmt.Printf("%sKernel Arch:%s \t\t%s\n", ColorGreen, ColorNone, host.KernelArch)
	fmt.Printf("%sKernel Version:%s \t%s\n", ColorGreen, ColorNone, host.KernelVersion)
	fmt.Printf("%sOS:%s \t\t\t%s\n", ColorGreen, ColorNone, host.OS)
	fmt.Printf("%sPlatform Family:%s \t%s\n", ColorGreen, ColorNone, host.PlatformFamily)
	// Convert uptime from seconds to a time.Duration
	uptimeDuration := time.Duration(host.Uptime) * time.Second
	fmt.Printf("%sUptime:%s \t\t%v\n", ColorGreen, ColorNone, uptimeDuration)
}

func MemoryInfo() {

	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting host information:", err)
		return
	}
	fmt.Printf("%sMemory Total:%s \t\t%v GB\n", ColorGreen, ColorNone, bytesToGigabytes(float64(memory.Total)))
	fmt.Printf("%sMemory Used:%s \t\t%v GB\n", ColorGreen, ColorNone, bytesToGigabytes(float64(memory.Used)))
	fmt.Printf("%sMemory Free:%s \t\t%v GB\n", ColorGreen, ColorNone, bytesToGigabytes(float64(memory.Free)))
}

func bytesToGigabytes(bytes float64) float64 {
	gigabytes := bytes / (1024 * 1024 * 1024)
	return math.Ceil(gigabytes)
}
