package logic

import (
	"context"
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
)

func (c *ConfigServerImpl) CreateEnv(ctx context.Context, config *config_writer.Config, name string) (*config_writer.Config, error) {
	panic("implement me")
}

func (c *ConfigServerImpl) UpdateEnv(ctx context.Context, env *config_writer.ConfigEnv) error {
	panic("implement me")
}
