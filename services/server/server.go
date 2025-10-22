package server

import (
	"sync"
	"ybub/services/conf"
	"ybub/services/emitter"
)

type SSHService struct {
	Config    *conf.ConfigManager
	Emitter   *emitter.Emitter
	cancelMap sync.Map
}

func New(cfg *conf.ConfigManager, e *emitter.Emitter) *SSHService {
	return &SSHService{
		Config:  cfg,
		Emitter: e,
	}
}
