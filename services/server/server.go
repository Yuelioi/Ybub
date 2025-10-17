package server

import (
	"ybub/services/conf"
	"ybub/services/emitter"
)

type SSHService struct {
	Config  *conf.ConfigManager
	Emitter *emitter.Emitter
}

func New(cfg *conf.ConfigManager, e *emitter.Emitter) *SSHService {
	return &SSHService{
		Config:  cfg,
		Emitter: e,
	}
}
