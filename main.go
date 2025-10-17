package main

import (
	"embed"
	"log"
	"ybub/app"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var iconFS embed.FS

func main() {
	// 配置应用
	config := app.Config{
		Name:        "ybub",
		Description: "Docker Manager",
		ConfigPath:  "ybub.yaml",
		Assets:      assets,
		IconFS:      iconFS,
	}

	// 创建应用实例
	application := app.New(config)

	// 运行应用
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}

}
