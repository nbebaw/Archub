package lib

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// CpuInfo displays CPU information including vendor, cores, frequency, and model name.
// Uses the gopsutil library to retrieve CPU details.
func CpuInfo() {
	// Get CPU information from the system
	cpu, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting host information:", err)
		return
	}

	// Display CPU vendor (e.g., GenuineIntel, AuthenticAMD)
	fmt.Printf("%sVendor ID:%s \t\t%s\n", ColorGreen, ColorNone, cpu[0].VendorID)

	// Calculate total number of cores across all CPUs
	var cores int
	for _, s := range cpu {
		cores = int(s.Cores) + cores
	}
	fmt.Printf("%sCores:%s \t\t\t%v\n", ColorGreen, ColorNone, cores)

	// Display CPU frequency in MHz
	fmt.Printf("%sMhz:%s \t\t\t%v\n", ColorGreen, ColorNone, cpu[0].Mhz)

	// Display CPU model name (e.g., Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz)
	fmt.Printf("%sModel Name:%s \t\t%v\n", ColorGreen, ColorNone, cpu[0].ModelName)
}

// HostInfo displays system host information including hostname, kernel, OS, and uptime.
// Uses the gopsutil library to retrieve host details.
func HostInfo() {
	// Get host information from the system
	host, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host information:", err)
		return
	}

	// Display system hostname
	fmt.Printf("%sHostname:%s \t\t%s\n", ColorGreen, ColorNone, host.Hostname)

	// Display kernel architecture (e.g., x86_64, aarch64)
	fmt.Printf("%sKernel Arch:%s \t\t%s\n", ColorGreen, ColorNone, host.KernelArch)

	// Display kernel version (e.g., 6.1.0-arch1-1)
	fmt.Printf("%sKernel Version:%s \t%s\n", ColorGreen, ColorNone, host.KernelVersion)

	// Display operating system name (e.g., linux, darwin)
	fmt.Printf("%sOS:%s \t\t\t%s\n", ColorGreen, ColorNone, host.OS)

	// Display platform family (e.g., arch, debian, rhel)
	fmt.Printf("%sPlatform Family:%s \t%s\n", ColorGreen, ColorNone, host.PlatformFamily)

	// Convert uptime from seconds to a human-readable duration
	uptimeDuration := time.Duration(host.Uptime) * time.Second
	fmt.Printf("%sUptime:%s \t\t%v\n", ColorGreen, ColorNone, uptimeDuration)
}

// MemoryInfo displays system memory information including total, used, and free memory.
// Uses the gopsutil library to retrieve memory statistics.
func MemoryInfo() {
	// Get virtual memory information from the system
	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting host information:", err)
		return
	}

	// Display memory statistics in GB
	fmt.Printf("%sMemory Total:%s \t\t%v GB\n", ColorGreen, ColorNone, bytesToGigabytes(float64(memory.Total)))
	fmt.Printf("%sMemory Used:%s \t\t%v GB\n", ColorGreen, ColorNone, bytesToGigabytes(float64(memory.Used)))
	fmt.Printf("%sMemory Free:%s \t\t%v GB\n", ColorGreen, ColorNone, bytesToGigabytes(float64(memory.Free)))
}

// bytesToGigabytes converts bytes to gigabytes and rounds up to the nearest integer.
//
// Parameters:
//   - bytes: Size in bytes
//
// Returns:
//   - float64: Size in gigabytes, rounded up
func bytesToGigabytes(bytes float64) float64 {
	// Convert bytes to GB (1 GB = 1024^3 bytes)
	gigabytes := bytes / (1024 * 1024 * 1024)
	// Round up to nearest integer
	return math.Ceil(gigabytes)
}
