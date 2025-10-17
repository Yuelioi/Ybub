package conf

import (
	"errors"
	"sync"
)

// 事务管理功能
type Transaction struct {
	inTransaction bool
	mu            sync.RWMutex
}

// 开始事务
func (c *ConfigManager) Begin() {
	c.transaction.mu.Lock()
	defer c.transaction.mu.Unlock()

	c.transaction.inTransaction = true
	c.logger.Debug().Msg("开始事务")
}

// 提交事务并保存
func (c *ConfigManager) Commit() error {
	c.transaction.mu.Lock()
	defer c.transaction.mu.Unlock()

	if !c.transaction.inTransaction {
		c.logger.Warn().Msg("未开始事务")
		return errors.New("未开始事务")
	}

	c.transaction.inTransaction = false
	c.logger.Debug().Msg("提交事务")
	return c.Save()
}

// 回滚事务(重新加载配置) todo 需要重新发送配置给前端
func (c *ConfigManager) Rollback() error {
	c.transaction.mu.Lock()
	defer c.transaction.mu.Unlock()

	if !c.transaction.inTransaction {
		c.logger.Warn().Msg("未开始事务")
		return errors.New("未开始事务")
	}

	c.transaction.inTransaction = false
	c.logger.Debug().Msg("回滚事务")

	// 重新加载配置文件
	newCfg, err := New(c.path)
	if err != nil {
		c.logger.Error().Err(err).Msg("回滚失败:无法重新加载配置")
		return err
	}

	// 恢复数据
	c.Servers = newCfg.Servers
	c.Projects = newCfg.Projects
	c.BackupDir = newCfg.BackupDir

	c.logger.Info().Msg("事务回滚成功")
	return nil
}

// 检查是否在事务中
func (c *ConfigManager) InTransaction() bool {
	c.transaction.mu.RLock()
	defer c.transaction.mu.RUnlock()
	return c.transaction.inTransaction
}

// save 内部保存方法,支持事务
func (c *ConfigManager) save() error {
	c.transaction.mu.RLock()
	defer c.transaction.mu.RUnlock()

	if c.transaction.inTransaction {
		c.logger.Debug().Msg("事务进行中,跳过保存")
		return nil
	}

	return c.Save()
}
