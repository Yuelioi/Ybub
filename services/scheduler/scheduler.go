package scheduler

import (
	"ybub/models"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type SchedulerService struct {
	Cron        *cron.Cron
	taskToEntry map[string]cron.EntryID
}

func New() *SchedulerService {
	return &SchedulerService{
		Cron:        cron.New(),
		taskToEntry: make(map[string]cron.EntryID),
	}
}

func (s *SchedulerService) AddTask(
	project models.Project,
	f func(project models.Project) error,
) error {
	p := project
	entryID, err := s.Cron.AddFunc(p.Schedule, func() {
		f(p)
	})
	if err != nil {
		return err
	}
	s.taskToEntry[p.ID] = entryID
	return nil
}

func (s *SchedulerService) RemoveTask(projectID string) {
	if entryID, ok := s.taskToEntry[projectID]; ok {
		s.Cron.Remove(entryID)
		delete(s.taskToEntry, projectID)
	}
}

func (s *SchedulerService) Start() {
	log.Info().Msg("定时任务已启动")
	s.Cron.Start()
}
