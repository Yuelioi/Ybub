package conf

import (
	"encoding/json"
	"errors"
	"os"
	"ybub/models"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.yaml.in/yaml/v4"
)

type ConfigManager struct {
	path        string           `json:"-" yaml:"-"`                   // 配置路径
	logger      zerolog.Logger   `json:"-" yaml:"-"`                   // 日志记录器
	transaction Transaction      `json:"-" yaml:"-"`                   // 事务管理
	BackupDir   string           `json:"backup_dir" yaml:"backup_dir"` // 备份储存路径
	Servers     []models.Server  `json:"servers" yaml:"servers"`
	Projects    []models.Project `json:"projects" yaml:"projects"`
}

func (c *ConfigManager) Save() error {
	c.logger.Debug().Str("path", c.path).Msg("开始保存配置文件")

	data, err := yaml.Marshal(c)
	if err != nil {
		c.logger.Error().Err(err).Msg("配置序列化失败")
		return err
	}

	if err := os.WriteFile(c.path, data, 0644); err != nil {
		c.logger.Error().Err(err).Str("path", c.path).Msg("写入配置文件失败")
		return err
	}

	c.logger.Info().Str("path", c.path).Msg("配置文件保存成功")
	return nil
}

// 从文件读取配置
func New(path string) (*ConfigManager, error) {
	logger := log.With().Str("component", "ConfigManager").Logger()
	logger.Info().Str("path", path).Msg("初始化配置管理器")

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		// 文件不存在：创建默认配置
		logger.Warn().Str("path", path).Msg("配置文件不存在,创建默认配置")

		cfg := &ConfigManager{
			Servers:  []models.Server{},
			Projects: []models.Project{},
			logger:   logger,
		}
		cfg.BackupDir = "backups"
		cfg.path = path

		if err := cfg.Save(); err != nil {
			logger.Error().Err(err).Msg("保存默认配置失败")
			return nil, err
		}

		logger.Info().Msg("默认配置创建成功")
		return cfg, nil
	} else if err != nil {
		logger.Error().Err(err).Str("path", path).Msg("读取配置文件失败")
		return nil, err
	}

	var cfg ConfigManager
	// 尝试 YAML 解析
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		logger.Warn().Err(err).Msg("YAML 解析失败,尝试 JSON 解析")
		// 如果 YAML 解析失败，尝试 JSON
		if err := json.Unmarshal(data, &cfg); err != nil {
			logger.Error().Err(err).Msg("配置文件解析失败(YAML 和 JSON 均失败)")
			return nil, err
		}
		logger.Info().Msg("使用 JSON 格式解析配置成功")
	} else {
		logger.Info().Msg("使用 YAML 格式解析配置成功")
	}

	cfg.path = path
	cfg.logger = logger

	logger.Info().
		Int("servers", len(cfg.Servers)).
		Int("projects", len(cfg.Projects)).
		Msg("配置加载完成")

	return &cfg, nil
}
