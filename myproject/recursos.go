package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func getCpuUsage() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(percent[0]))
}

func getMemoryUsage() int {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(memory.UsedPercent))
}

func getData() string {
	cpuData := "Cpu: " + strconv.Itoa(getCpuUsage()) + "% "
	memoryData := "Mem: " + strconv.Itoa(getMemoryUsage()) + "% "
	fmt.Println(cpuData + memoryData)
	return cpuData + memoryData
}

func onReady() {
	go func() {
		var result string
		for {
			result = getData()
			systray.SetTitle(result)
		}
	}()
}

func onExit() {

}

func ObtenerInfo() {
	systray.Run(onReady, onExit)
}
