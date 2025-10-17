package models

import "time"

type Event string

const (
	EventBackupProgress Event = "backup-progress"
	EventBackupComplete Event = "backup-complete"
)

const (
	EventSshOutput   Event = "ssh-output"
	EventSshComplete Event = "ssh-command-complete"
)

type SshCommandOutput struct {
	CommandID   string    `json:"commandId"`
	Status      Status    `json:"status"`
	Line        string    `json:"line"`
	CreatedTime time.Time `json:"createdTime"`
}

type BackupOutput struct {
	ProjectID string `json:"projectId"`
	Status    Status `json:"status"`
	Message   string `json:"message"`
}

type BackupProgress struct {
	ProjectID string `json:"projectId"`
	Status    Status `json:"status"`
	File      string `json:"file"`
	Size      int64  `json:"size"`
	Message   string `json:"message"`
}

type Status string

const (
	// 通用状态
	StatusPending     Status = "pending"
	StatusStarted     Status = "started"
	StatusRunning     Status = "running"
	StatusProgressing Status = "progressing"
	StatusSuccess     Status = "success"
	StatusFailed      Status = "failed"
	StatusCancelled   Status = "cancelled"
	StatusTimeout     Status = "timeout"
	StatusSkipped     Status = "skipped"

	// 备份任务特有状态
	StatusUploading Status = "uploading"
	StatusVerifying Status = "verifying"
	StatusRestoring Status = "restoring"
	StatusCompleted Status = "completed"
)
