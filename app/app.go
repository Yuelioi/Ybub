package app

import (
	"embed"
	"log/slog"
	"ybub/services/conf"
	"ybub/services/emitter"
	"ybub/services/scheduler"
	"ybub/services/server"
	servicecenter "ybub/services/service_center"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	app        *application.App
	mainWindow *application.WebviewWindow
	systray    *application.SystemTray
	log        zerolog.Logger
}

type Config struct {
	Name        string
	Description string
	ConfigPath  string
	Assets      embed.FS
	IconFS      embed.FS
}

func New(config Config) *App {
	logger := initLogger()
	log.Info().
		Str("appName", config.Name).
		Str("description", config.Description).
		Str("configPath", config.ConfigPath).
		Msg("开始初始化应用")

	// 创建Wails应用
	log.Debug().Msg("创建Wails应用实例")
	wailsApp := application.New(application.Options{
		Name:        config.Name,
		Description: config.Description,
		// Services:    createServices(config.ConfigPath),
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(config.Assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		SingleInstance: createSingleInstanceOptions(),
		LogLevel:       slog.LevelError,
	})
	log.Debug().Msg("Wails应用实例创建成功")

	configService, err := conf.New(config.ConfigPath)
	if err != nil {
		panic(err)
	}

	schedulerService := scheduler.New()
	emitterService := emitter.New(wailsApp)
	serverService := server.New(configService, emitterService)
	serviceCenter := servicecenter.New(configService, serverService, schedulerService)

	wailsApp.RegisterService(application.NewService(emitterService))
	wailsApp.RegisterService(application.NewService(configService))
	wailsApp.RegisterService(application.NewService(serverService))
	wailsApp.RegisterService(application.NewService(serviceCenter))

	for _, project := range configService.Projects {
		schedulerService.AddTask(project, serverService.BackupProjectData)
	}
	schedulerService.Start()

	app := &App{
		app: wailsApp,
		log: logger,
	}

	// 初始化各个组件
	log.Info().Msg("开始初始化应用组件")

	log.Debug().Msg("设置上下文菜单")
	app.setupContextMenu()

	log.Debug().Msg("设置主窗口")
	app.setupMainWindow()

	log.Debug().Msg("设置系统托盘")
	app.setupSystemTray(config.IconFS)

	log.Debug().Msg("启动事件发射器")
	app.startEventEmitter()

	log.Info().Msg("应用初始化完成")
	return app
}

func (a *App) Run() error {
	a.log.Info().Msg("开始运行应用")

	err := a.app.Run()
	if err != nil {
		a.log.Error().Err(err).Msg("应用运行失败")
		return err
	}

	// 居中窗口
	a.log.Debug().Msg("居中显示主窗口")
	a.mainWindow.Center()

	a.log.Info().Msg("应用运行成功")
	return nil
}

func (a *App) Quit() {
	a.log.Info().Msg("退出应用")
	a.app.Quit()
	a.log.Info().Msg("应用已退出")
}

func (a *App) GetApp() *application.App {
	a.log.Debug().Msg("获取Wails应用实例")
	return a.app
}

func (a *App) GetMainWindow() *application.WebviewWindow {
	a.log.Debug().Msg("获取主窗口实例")
	return a.mainWindow
}
