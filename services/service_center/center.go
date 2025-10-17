package servicecenter

import (
	"fmt"
	"time"
	"ybub/models"
	"ybub/services/conf"
	"ybub/services/scheduler"
	"ybub/services/server"

	"github.com/robfig/cron/v3"
)

type ServerCenter struct {
	conf      *conf.ConfigManager
	ssh       *server.SSHService
	scheduler *scheduler.SchedulerService
}

func New(conf *conf.ConfigManager,
	ssh *server.SSHService,
	scheduler *scheduler.SchedulerService) *ServerCenter {

	return &ServerCenter{
		conf:      conf,
		ssh:       ssh,
		scheduler: scheduler,
	}
}

func (sc *ServerCenter) CreateProjectWithScheduler(p models.Project) error {

	if p.EnableBackup {
		sc.scheduler.AddTask(p, sc.ssh.BackupProjectData)
	}
	return sc.conf.CreateProject(p)
}

func (sc *ServerCenter) RemoveServer(id string) error {

	sc.conf.Begin()
	projects := sc.conf.ListProjectsByServerID(id)

	for _, p := range projects {
		if err := sc.conf.RemoveProject(p.ID); err != nil {
			return err
		}
	}
	if err := sc.conf.RemoveServer(id); err != nil {
		return err
	}

	return sc.conf.Commit()
}

func (sc *ServerCenter) RemoveProject(id string) (err error) {
	if err = sc.conf.RemoveProject(id); err != nil {
		return err
	}
	sc.scheduler.RemoveTask(id)
	return
}

func (sc *ServerCenter) UpdateProject(p models.Project) error {
	oldProject, err := sc.conf.GetProject(p.ID)
	if err != nil {
		return err
	}

	if oldProject.EnableBackup && !p.EnableBackup {
		sc.scheduler.RemoveTask(p.ID)
	}

	if p.EnableBackup && !oldProject.EnableBackup {
		sc.scheduler.AddTask(p, sc.ssh.BackupProjectData)
	}

	if (p.EnableBackup && oldProject.EnableBackup) && p.Schedule != oldProject.Schedule {
		sc.scheduler.RemoveTask(p.ID)
		sc.scheduler.AddTask(p, sc.ssh.BackupProjectData)
	}

	sc.conf.UpdateProject(p)

	return nil
}

func (sc *ServerCenter) BackupProjectManually(p models.Project) error {
	err := sc.ssh.BackupProjectData(p)
	if err != nil {
		return err
	}

	if p.Schedule != "" {
		schedule, err := cron.ParseStandard(p.Schedule)
		if err != nil {
			return fmt.Errorf("解析 cron 表达式失败: %w", err)
		}
		*p.NextBackup = schedule.Next(time.Now())
	} else {
		// 默认加一天
		*p.NextBackup = time.Now().Add(24 * time.Hour)
	}

	*p.LastBackup = time.Now()
	return sc.conf.UpdateProject(p)
}
