package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type Response struct {
	Status string
	Message string
	Data interface{}
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	textTable := NewTextEditor()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "gim-wails",
		Width:  1024,
		Height: 768,
		OnStartup: app.startup,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			app,
			textTable,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
