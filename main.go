package main

import (
	"embed"
	"fmt"

	"colintester.com/ergo-sum/go/IPC"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:src/dist
var assets embed.FS

//go:embed build/appicon.png
var appIcon []byte

func main() {
	service := IPC.NewService()

	// Specific macOS desktop app options.
	macOptions := &mac.Options{
		Appearance: mac.NSAppearanceNameDarkAqua,
		About: &mac.AboutInfo{
			Title:   "ErgoSum",           // substitute your own App's name.
			Message: "© 2023 Cogito Inc", // subtitute your own ‘Project Owner’.
			Icon:    appIcon,             // here we reference our embedded custom icon.
		},
	}

	// Define Wails operating options:
	wailsOptions := &options.App{
		Title:  "ErgoSum", // substitute your own App's name.
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// Bind our Service instance for IPC.
		Bind: []interface{}{
			service,
		},
		Mac: macOptions,
	}

	// Create application with defined options:
	if err := wails.Run(wailsOptions); err != nil {
		fmt.Println("Wails setup error:", err.Error())
	}
}
