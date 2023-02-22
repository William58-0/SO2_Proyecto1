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
	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math"
	"strconv"
	"time"
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
	/*
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quits this app")
		go func() {
			for {
				select {
				case <-mQuit.ClickedCh:
					systray.Quit()
					return
				}

			}

		}()
	*/
}

func onExit() {

}

func main() {
	systray.Run(onReady, onExit)
}
