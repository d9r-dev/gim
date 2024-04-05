package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	filePath string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenFile() {
	response := Response{}
	options := runtime.OpenDialogOptions{
		Title: "Hello",
	}

	filePath, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil {
		log.Printf("Error retrieving file path. %v", err)
		return
	} else if filePath == "" {
		return
	}

	a.filePath = filePath
	runtime.WindowSetTitle(a.ctx, "gim - "+filepath.Base(filePath))

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file. %v", err)
		return
	}

	response.Status = "success"
	response.Message = filePath
	response.Data = string(data)

	runtime.EventsEmit(a.ctx, "onFileRead", response)
}
