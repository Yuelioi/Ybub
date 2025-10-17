package app

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// https://v3alpha.wails.io/learn/systray/

// 设置系统托盘
func (a *App) setupSystemTray(iconFS embed.FS) {
	iconBytes, err := iconFS.ReadFile("build/appicon.png")
	if err != nil {
		log.Printf("Failed to read icon file: %v", err)
		return
	}

	systray := a.app.SystemTray.New()
	// systray.SetLabel("YBub")
	systray.SetIcon(iconBytes)

	// 设置点击事件
	systray.OnClick(func() {
		if a.mainWindow != nil {
			a.mainWindow.Focus()
			a.mainWindow.UnMinimise()
		}
	})
	a.systray = systray

	// 创建托盘菜单
	a.createTrayMenu()
}

// 创建托盘菜单
func (a *App) createTrayMenu() {
	menu := application.NewMenu()

	// 退出菜单项
	exitItem := menu.Add("退出")
	exitItem.SetTooltip("关闭软件")
	exitItem.OnClick(func(*application.Context) {
		a.app.Quit()
	})

	// 显示/隐藏菜单项
	showItem := menu.Add("显示/隐藏")
	showItem.SetTooltip("显示或隐藏主窗口")
	showItem.OnClick(func(*application.Context) {
		if a.mainWindow != nil {
			a.mainWindow.Focus()
			a.mainWindow.Show()
			a.mainWindow.UnMinimise()
		}
	})

	a.systray.SetMenu(menu)
}

// 更新托盘图标
func (a *App) UpdateTrayIcon(iconData []byte) {
	if a.systray != nil {
		a.systray.SetIcon(iconData)
	}
}

// 更新托盘标签
func (a *App) UpdateTrayLabel(label string) {
	if a.systray != nil {
		a.systray.SetLabel(label)
	}
}
