package cfg

import (
	"context"
	"github.com/tars-vcms/config-writer/entity/config"
	"gorm.io/gorm"
)

const (
	// OptionAll 生成随机字符串代表搜索所有
	OptionAll = "IDLyocviS9nPh6bW"
)

const (
	EnvColumn  = "Envs"
	ItemColumn = "Items"
)

type configOption struct {
	Column string
	Value  interface{}
}

type ConfigOption interface {
	apply() func(db *gorm.DB) *gorm.DB
}

type Config interface {
	GetByIDs(ids []uint64, opts ...ConfigOption) (*[]config.Config, error)

	GetByNames(names []string, opts ...ConfigOption) (*[]config.Config, error)

	InputConfig(ctx context.Context, cfg *config.Config) error

	SaveConfig(ctx context.Context, cfg *config.Config) error
}

func New(db *gorm.DB) Config {
	return newConfigImpl(db)
}

func newConfigOption(f func(*configOption)) ConfigOption {
	return newConfigOptionImpl(f)
}

func WithEnv(name string) ConfigOption {
	return newConfigOption(func(o *configOption) {
		o.Column = EnvColumn
		o.Value = name
	})
}

func WithItem(name string) ConfigOption {
	return newConfigOption(func(o *configOption) {
		o.Column = ItemColumn
		o.Value = name
	})
}
