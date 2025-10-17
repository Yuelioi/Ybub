package models

import "time"

type ProjectStatus string

const (
	ProjectStatusRunning       ProjectStatus = "running"
	ProjectStatusStopped       ProjectStatus = "stopped"
	ProjectStatusError         ProjectStatus = "error"
	ProjectBackupStatusRunning ProjectStatus = "running"
	ProjectBackupStatusStopped ProjectStatus = "stopped"
	ProjectBackupStatusError   ProjectStatus = "error"
)

type Project struct {
	ID       string        `json:"id" yaml:"id"`
	ServerID string        `json:"serverId" yaml:"serverId"`
	Name     string        `json:"name" yaml:"name"`
	Path     string        `json:"path" yaml:"path"`
	Status   ProjectStatus `json:"status" yaml:"status"`
	DataPath string        `json:"dataPath,omitempty" yaml:"dataPath,omitempty"`

	EnableBackup bool          `json:"enableBackup" yaml:"enableBackup"`
	Schedule     string        `json:"schedule" yaml:"schedule"`
	BackupStatus ProjectStatus `json:"backupStatus" yaml:"backupStatus"`
	LastBackup   *time.Time    `json:"lastBackup,omitempty" yaml:"lastBackup,omitempty"`
	NextBackup   *time.Time    `json:"nextBackup,omitempty" yaml:"nextBackup,omitempty"`
}
