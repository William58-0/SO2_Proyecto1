package main

import (
	"syscall"
	"strconv"
	"fmt"
)

// FUENTE: https://github.com/pbnjay/memory/blob/7b4eea64cf58/memory_linux.go
func sysTotalMemory64() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}

	return uint64(in.Totalram) * uint64(in.Unit)
}

func sysFreeMemory64() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}

	return uint64(in.Freeram) * uint64(in.Unit)
}

// ----------------------------- Para 32 bits
func sysTotalMemory32() uint32 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}

	return uint32(in.Totalram) * uint32(in.Unit)
}

func sysFreeMemory32() uint32 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}

	return uint32(in.Freeram) * uint32(in.Unit)
}

//------------------------------------------

func main() {
	factorGb := 1024 * 1024 * 1024;

	// FUENTE: https://programming-idioms.org/idiom/160/detect-if-32-bit-or-64-bit-architecture/1981/go
	if (strconv.IntSize==64) {
		libre := sysFreeMemory64();
		total := sysTotalMemory64();

		fmt.Println(libre);
		fmt.Println(total);

		libre1 := fmt.Sprintf("%.2f", (float64(libre) / float64(factorGb)))
		total1 := fmt.Sprintf("%.2f", (float64(total) / float64(factorGb)))

		fmt.Println(libre1);
		fmt.Println(total1);
		
	} else if(strconv.IntSize==32) {
		libre := sysFreeMemory32();
		total := sysTotalMemory32();

		fmt.Println(libre);
		fmt.Println(total);

		libre1 := fmt.Sprintf("%.2f", (float64(libre) / float64(factorGb)))
		total1 := fmt.Sprintf("%.2f", (float64(total) / float64(factorGb)))

		fmt.Println(libre1);
		fmt.Println(total1);
	}

	
}