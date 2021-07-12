package logic

import (
	"context"
	"github.com/tars-vcms/config-writer/entity/configenv"
	"github.com/tars-vcms/config-writer/entity/configitem"
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
)

func (c *ConfigServerImpl) PublishItem(ctx context.Context, env *configenv.ConfigEnv, item *config_writer.ConfigItem) error {
	panic("implement me")
}

func (c *ConfigServerImpl) SetActiveItem(ctx context.Context, item *configitem.ConfigItem, version uint64) error {
	panic("implement me")
}
