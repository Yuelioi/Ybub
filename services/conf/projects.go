package conf

import (
	"errors"
	"ybub/models"
)

func (c *ConfigManager) ListProjects() []models.Project {
	c.logger.Debug().Int("count", len(c.Projects)).Msg("获取项目列表")
	return c.Projects
}

func (c *ConfigManager) ListProjectsByServerID(serverID string) []models.Project {
	var result []models.Project

	for _, p := range c.Projects {
		if p.ServerID == serverID {
			result = append(result, p)
		}
	}

	c.logger.Debug().
		Str("server_id", serverID).
		Int("count", len(result)).
		Msg("获取项目列表")

	return result
}

// 前端请勿调用
func (c *ConfigManager) CreateProject(p models.Project) error {
	c.logger.Debug().Str("project_id", p.ID).Str("name", p.Name).Msg("添加项目")

	if _, found := findIndexByID(c.Projects, p.ID, func(v models.Project) string { return v.ID }); found {
		c.logger.Warn().Str("project_id", p.ID).Msg("项目 ID 已存在")
		return errors.New("项目 ID 已存在")
	}

	c.Projects = append(c.Projects, p)
	if err := c.save(); err != nil {
		return err
	}

	c.logger.Info().Str("project_id", p.ID).Str("name", p.Name).Msg("项目添加成功")
	return nil
}

// 前端请勿调用
func (c *ConfigManager) UpdateProject(p models.Project) error {
	c.logger.Debug().Str("project_id", p.ID).Msg("更新项目")

	if idx, found := findIndexByID(c.Projects, p.ID, func(v models.Project) string { return v.ID }); found {
		c.Projects[idx] = p
		if err := c.save(); err != nil {
			return err
		}
		c.logger.Info().Str("project_id", p.ID).Str("name", p.Name).Msg("项目更新成功")
		return nil
	}

	c.logger.Warn().Str("project_id", p.ID).Msg("项目不存在")
	return errors.New("项目不存在")
}

// 前端请勿调用
func (c *ConfigManager) RemoveProject(id string) error {
	c.logger.Debug().Str("project_id", id).Msg("删除项目")

	if idx, found := findIndexByID(c.Projects, id, func(v models.Project) string { return v.ID }); found {
		c.Projects = append(c.Projects[:idx], c.Projects[idx+1:]...)

		if err := c.save(); err != nil {
			return err
		}

		c.logger.Info().
			Str("project_id", id).
			Msg("项目删除成功")
		return nil
	}

	c.logger.Warn().Str("project_id", id).Msg("项目不存在")
	return errors.New("项目不存在")
}

func (c *ConfigManager) GetProject(id string) (*models.Project, error) {
	if idx, found := findIndexByID(c.Projects, id, func(v models.Project) string { return v.ID }); found {
		c.logger.Debug().Str("project_id", id).Msg("获取项目信息")
		return &c.Projects[idx], nil
	}
	c.logger.Warn().Str("project_id", id).Msg("项目不存在")
	return nil, errors.New("项目不存在")
}
