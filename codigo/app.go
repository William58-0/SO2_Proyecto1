package main

import (
	"context"
	"fmt"
	"time"

	//"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"encoding/json"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go func() {
		for {
			resources := getRecursos()
			json, err := json.Marshal(resources)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return;
			}
			fmt.Println(string(json))

			runtime.EventsEmit(ctx, "recursos",string(json))
			time.Sleep(1 * time.Second)
		}
	}()

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) WailsInit(runtime *wails.Runtime) error {
	// s.log = runtime.Log.New("Stats")

	

	return nil
}