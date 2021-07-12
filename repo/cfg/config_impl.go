package cfg

import (
	"context"
	"github.com/tars-vcms/config-writer/entity/config"
	"gorm.io/gorm"
)

type ConfigImpl struct {
	db *gorm.DB
}

func (c *ConfigImpl) SaveConfig(ctx context.Context, cfg *config.Config) error {
	tx := c.db.WithContext(ctx)
	result := tx.Save(cfg)
	return result.Error
}

func newConfigImpl(db *gorm.DB) *ConfigImpl {
	return &ConfigImpl{db: db}
}

func (c *ConfigImpl) GetByIDs(ids []uint64, opts ...ConfigOption) (*[]config.Config, error) {
	tx := c.injectOption(opts, c.db)
	var cfgEntities []config.Config
	isFirst := true
	for _, value := range ids {
		if isFirst {
			tx.Where("id = ?", value)
			isFirst = false
		} else {
			tx.Or("id = ?", value)
		}
	}
	tx = c.db.Find(&cfgEntities)
	return &cfgEntities, tx.Error
}

func (c *ConfigImpl) GetByNames(names []string, opts ...ConfigOption) (*[]config.Config, error) {
	tx := c.injectOption(opts, c.db)
	var cfgEntities []config.Config
	isFirst := true
	for _, value := range names {
		if isFirst {
			tx.Where("name = ?", value)
			isFirst = false
		} else {
			tx.Or("name = ?", value)
		}
	}
	tx = c.db.Find(&cfgEntities)
	return &cfgEntities, tx.Error
}

func (c *ConfigImpl) injectOption(opts []ConfigOption, tx *gorm.DB) *gorm.DB {
	for _, opt := range opts {
		tx.Scopes(opt.apply())
	}
	return tx
}

func (c *ConfigImpl) InputConfig(ctx context.Context, cfg *config.Config) error {
	tx := c.db.WithContext(ctx)
	result := tx.Create(cfg)
	return result.Error
}
