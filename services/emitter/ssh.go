package emitter

import (
	"time"
	"ybub/models"
)

func (e *Emitter) EmitSshOutput(commandID string, status models.Status, line string) {
	e.Emit(models.EventSshOutput, models.SshCommandOutput{
		CommandID:   commandID,
		Status:      status,
		Line:        line,
		CreatedTime: time.Now(),
	})
}

func (e *Emitter) EmitSshCompleted(commandID string) {
	e.Emit(models.EventSshComplete, models.SshCommandOutput{
		CommandID:   commandID,
		Status:      models.StatusCompleted,
		CreatedTime: time.Now(),
	})
}
