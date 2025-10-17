package app

import (
	"time"
)

// 事件发射器示例
func (a *App) startEventEmitter() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			now := time.Now().Format(time.RFC1123)
			a.app.Event.Emit("time", now)
		}
	}()

}
