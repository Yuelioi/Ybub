package models

import "time"

type ServerStatus string

const (
	ServerStatusConnected    ServerStatus = "connected"
	ServerStatusDisconnected ServerStatus = "disconnected"
	ServerStatusConnecting   ServerStatus = "connecting"
)

type Server struct {
	ID            string       `json:"id" yaml:"id"`
	Name          string       `json:"name" yaml:"name"`
	Host          string       `json:"host" yaml:"host"`
	Port          int          `json:"port" yaml:"port"`
	Username      string       `json:"username" yaml:"username"`
	Password      string       `json:"password,omitempty" yaml:"password,omitempty"`
	IdentityFile  string       `json:"identityFile,omitempty" yaml:"identityFile,omitempty"`
	Status        ServerStatus `json:"status" yaml:"status"`
	LastConnected *time.Time   `json:"lastConnected,omitempty" yaml:"lastConnected,omitempty"`
}
