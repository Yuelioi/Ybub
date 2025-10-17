package emitter

import (
	"context"
	"ybub/models"

	"github.com/rs/zerolog/log"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Emitter struct {
	app *application.App
	ctx context.Context
}

func New(app *application.App) *Emitter {
	return &Emitter{app: app}
}

func (e *Emitter) ServiceStartup(ctx context.Context, _ application.ServiceOptions) error {
	e.ctx = ctx
	log.Info().Msg("SSHEmitter 启动完成")
	return nil
}

func (e *Emitter) Emit(event models.Event, data any) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Event emit panic [%s]: %v", event, r)
		}
	}()

	e.app.Event.Emit(string(event), data)
}

// 仅用于导出模型, 前端请勿使用
func (e *Emitter) Export(models.BackupOutput, models.BackupProgress, models.SshCommandOutput) {

}
