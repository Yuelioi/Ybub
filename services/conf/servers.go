package conf

import (
	"errors"
	"ybub/models"
)

func (c *ConfigManager) ListServers() []models.Server {
	c.logger.Debug().Int("count", len(c.Servers)).Msg("获取服务器列表")
	return c.Servers
}

func (c *ConfigManager) CreateServer(s models.Server) error {
	c.logger.Debug().Str("server_id", s.ID).Str("name", s.Name).Msg("添加服务器")

	if _, found := findIndexByID(c.Servers, s.ID, func(v models.Server) string { return v.ID }); found {
		c.logger.Warn().Str("server_id", s.ID).Msg("服务器 ID 已存在")
		return errors.New("服务器 ID 已存在")
	}

	c.Servers = append(c.Servers, s)
	if err := c.save(); err != nil {
		return err
	}

	c.logger.Info().Str("server_id", s.ID).Str("name", s.Name).Msg("服务器添加成功")
	return nil
}

func (c *ConfigManager) UpdateServer(id string, s models.Server) error {
	c.logger.Debug().Str("server_id", id).Msg("更新服务器")

	if idx, found := findIndexByID(c.Servers, id, func(v models.Server) string { return v.ID }); found {
		s.ID = id
		c.Servers[idx] = s
		if err := c.save(); err != nil {
			return err
		}
		c.logger.Info().Str("server_id", id).Str("name", s.Name).Msg("服务器更新成功")
		return nil
	}

	c.logger.Warn().Str("server_id", id).Msg("服务器不存在")
	return errors.New("服务器不存在")
}

// 前端请勿调用
func (c *ConfigManager) RemoveServer(id string) error {
	c.logger.Debug().Str("server_id", id).Msg("删除服务器")

	if idx, found := findIndexByID(c.Servers, id, func(v models.Server) string { return v.ID }); found {
		// 删除服务器
		c.Servers = append(c.Servers[:idx], c.Servers[idx+1:]...)

		if err := c.save(); err != nil {
			return err
		}

		c.logger.Info().
			Str("server_id", id).
			Msg("服务器删除成功")
		return nil
	}

	c.logger.Warn().Str("server_id", id).Msg("服务器不存在")
	return errors.New("服务器不存在")
}

func (c *ConfigManager) GetServer(id string) (*models.Server, error) {
	if idx, found := findIndexByID(c.Servers, id, func(v models.Server) string { return v.ID }); found {
		c.logger.Debug().Str("server_id", id).Msg("获取服务器信息")
		return &c.Servers[idx], nil
	}
	c.logger.Warn().Str("server_id", id).Msg("服务器不存在")
	return nil, errors.New("服务器不存在")
}
