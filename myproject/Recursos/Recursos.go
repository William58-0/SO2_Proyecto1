package recursos

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	
	human "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)


func getCpuUsage() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(percent[0]))
}

func getDiskUsage() {
	formatter := "%-14s %7s %7s %7s %4s %s\n"
	fmt.Printf(formatter, "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")

	parts, _ := disk.Partitions(true)
	
	for _, p := range parts {
		device := p.Mountpoint
		// solo funciona así en linux
		if device == "/" {
			s, _ := disk.Usage(device)

			if s.Total == 0 {
				continue
			}
	
			percent := fmt.Sprintf("%2.f%%", s.UsedPercent)
	
			return percent
		}
	}
}

func getData() string {
	cpuData := "Cpu: " + strconv.Itoa(getCpuUsage()) + "% "
	diskData := "Cpu: " + strconv.Itoa(getDiskUsage()) + "% "
	fmt.Println(cpuData + memoryData)
	return cpuData + diskData
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
