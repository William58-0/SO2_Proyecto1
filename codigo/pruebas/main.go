package main

import (
	"syscall"
 
	"fmt"
)

// FUENTE: https://github.com/pbnjay/memory/blob/7b4eea64cf58/memory_linux.go
func sysTotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}

func sysFreeMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Freeram) * uint64(in.Unit)
}

func main() {
	libre := sysFreeMemory();
	total := sysTotalMemory();

	factorGb := 1024 * 1024 * 1024;

	fmt.Println(libre);
	fmt.Println(total);

	libre1 := fmt.Sprintf("%.2f", (float64(libre) / float64(factorGb)))
	total1 := fmt.Sprintf("%.2f", (float64(total) / float64(factorGb)))

	fmt.Println(libre1);
	fmt.Println(total1);
}