
package main

import (
	"fmt"
	"runtime"

	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"

	// human "github.com/dustin/go-humanize"
)

// CPU struct
type CPU struct {
	Porcentaje int `json:"Porcentaje"`
	Nucleos int `json:"Nucleos"`
}

// Disco struct
type Disco struct {
	Porcentaje int `json:"Porcentaje"`
	Usado string `json:"Usado"`
	Disponible string `json:"Disponible"`
	Total string `json:"Total"`
}

// Recursos struct
type Recursos struct {
	CPU CPU `json:"CPU"`
	Disco Disco `json:"Disco"`
}

func getCPUInfo() CPU {
	// porcentaje de uso
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}

	porcentaje := int(math.Ceil(percent[0]))

	// nucleos del CPU
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	// struct CPU
	cpu := CPU{Porcentaje: porcentaje, Nucleos: cores}
	return cpu
}

func getDiskInfo() Disco {
	parts, _ := disk.Partitions(true)
	disco := Disco{}

	for _, p := range parts {
		device := p.Mountpoint
		// solo funciona así en linux
		if device == "/" {
			s, _ := disk.Usage(device)

			if s.Total == 0 {
				continue
			}
			
			porcentaje := int(math.Round((float64(s.Total - s.Free)*100.0/float64(s.Total))))
			libre := fmt.Sprintf("%.1f", (float64(s.Free) / 1000000000.0))
			total := fmt.Sprintf("%.1f", (float64(s.Total) / 1000000000.0))
			usado := fmt.Sprintf("%.1f", (float64(s.Total - s.Free) / 1000000000.0))

			disco = Disco{
				Porcentaje: porcentaje,
				Usado: usado,
				Disponible: libre,
				Total: total,
			}

			break
		}
	}
	return disco
}

func getRecursos() Recursos {
	recursos := Recursos{CPU: getCPUInfo(), Disco: getDiskInfo()}
	return recursos
}

func PruebaDisco() {
	fmt.Println("siuuu")
}
