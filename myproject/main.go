/*
package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
*/
package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	// "github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	
	// human "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)


func getCpuUsage() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	return int(math.Ceil(percent[0]))
}

func getDiskUsage() int{
	parts, _ := disk.Partitions(true)
	percent := 0
	for _, p := range parts {
		device := p.Mountpoint
		// solo funciona así en linux
		if device == "/" {
			s, _ := disk.Usage(device)

			if s.Total == 0 {
				continue
			}
	
			percent = int(math.Ceil(s.UsedPercent))
			break
		}
	}
	return percent
}

func getData() string {
	cpuData := "Cpu: " + strconv.Itoa(getCpuUsage()) + "% "
	diskData := "Disk: " + strconv.Itoa(getDiskUsage()) + "% "
	// fmt.Println(cpuData + diskData)
	return cpuData + diskData
}

func onReady() {
	var result string
	for {
		result = getData()
		fmt.Println(result)
		//systray.SetTitle(result)
	}
}

func main() {
	onReady()
	/*
	for {
		fmt.Println("hola")
	}
	*/
}
