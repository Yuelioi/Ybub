package app

import (
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// createSingleInstanceOptions 创建单实例选项
func createSingleInstanceOptions() *application.SingleInstanceOptions {
	return &application.SingleInstanceOptions{
		UniqueID: "com.ybub.yueli",
		OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
			log.Printf("Second instance launched with args: %v", data.Args)
			log.Printf("Working directory: %s", data.WorkingDir)
			log.Printf("Additional data: %v", data.AdditionalData)
		},
		AdditionalData: map[string]string{
			"launchtime": time.Now().String(),
		},
	}
}
