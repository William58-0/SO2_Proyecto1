package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
			fmt.Println("ejecutando")
			// ctx.Events.Emit("cpu_usage")
			//EventsEmit(ctx, "cpu_usage")
			runtime.EventsEmit(ctx, "cpu_usage","hola")
			time.Sleep(1 * time.Second)
		}
	}()

}


// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Greet %s, It's show time!", name)
}

func (a *App) Hola() string {
	return "Holaa, It's show time!"
}

// func Percent(interval time.Duration, percpu bool) ([]float64, error)

func (a *App) GetUsoCPU() string {

	return "Holaa, It's show time!"
}
