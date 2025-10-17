package app

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

//https://v3alpha.wails.io/learn/context-menu/

// 设置上下文菜单示例
func (a *App) setupContextMenu() {
	contextMenu := application.NewContextMenu("test")

	// 添加菜单项
	clickMe := contextMenu.Add("Click to show context data")
	dataLabel := contextMenu.Add("Current data: None")

	// 设置点击事件
	clickMe.OnClick(func(ctx *application.Context) {
		data := ctx.ContextMenuData()
		dataLabel.SetLabel("Current data: " + data)
		contextMenu.Update()
	})
}

// ContextMenuBuilder 上下文菜单构建器
type ContextMenuBuilder struct {
	menu *application.ContextMenu
}

// NewContextMenuBuilder 创建新的上下文菜单构建器
func NewContextMenuBuilder(name string) *ContextMenuBuilder {
	return &ContextMenuBuilder{
		menu: application.NewContextMenu(name),
	}
}

// AddItem 添加菜单项
func (cmb *ContextMenuBuilder) AddItem(text string, onClick func(*application.Context)) *ContextMenuBuilder {
	item := cmb.menu.Add(text)
	if onClick != nil {
		item.OnClick(onClick)
	}
	return cmb
}

// AddSeparator 添加分隔符
func (cmb *ContextMenuBuilder) AddSeparator() *ContextMenuBuilder {
	// Wails v3 中的分隔符添加方式可能需要根据实际API调整
	return cmb
}

// Build 构建菜单
func (cmb *ContextMenuBuilder) Build() *application.ContextMenu {
	return cmb.menu
}

// CreateDefaultContextMenu 创建默认的上下文菜单
func CreateDefaultContextMenu() *application.ContextMenu {
	builder := NewContextMenuBuilder("default")

	return builder.
		AddItem("复制", func(ctx *application.Context) {
			// 复制逻辑
		}).
		AddItem("粘贴", func(ctx *application.Context) {
			// 粘贴逻辑
		}).
		AddItem("刷新", func(ctx *application.Context) {
			// 刷新逻辑
		}).
		Build()
}
