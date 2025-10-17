package app

import (
	"os"

	"github.com/Yuelioi/gkit/logx/zero"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger() (logger zerolog.Logger) {
	isProduction := os.Getenv("PRODUCTION") == "true"

	if !isProduction {
		// 生产模式:使用文件日志
		logFile := &lumberjack.Logger{
			Filename:   "logs/app.log", // 日志文件路径
			MaxSize:    10,             // 每个文件最大 10 MB
			MaxBackups: 7,              // 最多保留 7 个旧文件
			MaxAge:     30,             // 保留 30 天
			Compress:   true,           // 是否压缩旧日志 (.gz)
		}

		logger := zerolog.New(logFile).With().Timestamp().Logger()
		log.Info().Msg("生产模式:日志输出到文件")

		return logger

	} else {
		// 开发模式:使用控制台日志
		logger = zero.Default()
		log.Logger = logger

		log.Info().Msg("开发模式:日志输出到控制台")
	}

	return

}
