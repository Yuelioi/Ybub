package emitter

import "ybub/models"

func (e *Emitter) EmitBackupTaskOutput(projectID string, status models.Status, message string) {
	e.Emit(models.EventBackupProgress, models.BackupOutput{
		ProjectID: projectID,
		Status:    status,
		Message:   message,
	})
}

func (e *Emitter) EmitBackupTaskProgress(projectID string, file string, size int64, message string) {
	e.Emit(models.EventBackupProgress, models.BackupProgress{
		ProjectID: projectID,
		Status:    models.StatusProgressing,
		Message:   message,
		File:      file,
		Size:      size,
	})
}

func (e *Emitter) EmitBackupTaskFinished(projectID string, message string) {
	e.Emit(models.EventBackupComplete, models.BackupOutput{
		ProjectID: projectID,
		Status:    models.StatusCompleted,
		Message:   message,
	})
}
