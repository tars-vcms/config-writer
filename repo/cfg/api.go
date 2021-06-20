package cfg

import (
	"config-server/entity/config"
	"context"
)

type Config interface {
	GetByID(id uint) (*config.Config, error)

	GetBySecret(secret string) (*config.Config, error)

	InputConfig(ctx context.Context, cfg *config.Config) error
}

func New() Config {
	return newConfigImpl()
}
