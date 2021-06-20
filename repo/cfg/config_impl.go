package cfg

import (
	"config-server/entity/config"
	"context"
)

type ConfigImpl struct {
}

func newConfigImpl() *ConfigImpl {
	return &ConfigImpl{}
}

func (c *ConfigImpl) GetByID(id uint) (*config.Config, error) {
	panic("implement me")
}

func (c *ConfigImpl) GetBySecret(secret string) (*config.Config, error) {
	panic("implement me")
}

func (c *ConfigImpl) InputConfig(ctx context.Context, cfg *config.Config) error {
	panic("implement me")
}
