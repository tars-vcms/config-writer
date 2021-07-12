package logic

import (
	"context"
	"github.com/tars-vcms/config-writer/entity/configenv"
	"github.com/tars-vcms/config-writer/entity/configitem"
	"github.com/tars-vcms/vcms-protocol/config/config-writer"
)

type ConfigServer interface {
	CreateConfig(ctx context.Context, name string) (*config_writer.Config, error)

	UpdateConfig(ctx context.Context, config *config_writer.Config) error

	CreateEnv(ctx context.Context, config *config_writer.Config, name string) (*config_writer.Config, error)

	UpdateEnv(ctx context.Context, env *config_writer.ConfigEnv) error

	GetConfigs(CIds []uint64, Names []string, withEnv bool) ([]*config_writer.Config, error)

	GetItems()

	PublishItem(ctx context.Context, env *configenv.ConfigEnv, item *config_writer.ConfigItem) error

	SetActiveItem(ctx context.Context, item *configitem.ConfigItem, version uint64) error
}

func NewConfigServer() ConfigServer {
	return newConfigServerImpl()
}
