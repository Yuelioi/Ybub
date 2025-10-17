package app

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

// setupMainWindow 设置主窗口
func (a *App) setupMainWindow() {
	a.mainWindow = a.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "YBUB",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		Width:            1366,
		Height:           768,
		MinWidth:         800,
		MinHeight:        600,
		Frameless:        true,
	})

	a.mainWindow.SetURL("/")
}
