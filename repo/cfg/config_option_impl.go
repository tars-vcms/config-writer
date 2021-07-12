package cfg

import "gorm.io/gorm"

func newConfigOptionImpl(f func(option *configOption)) *ConfigOptionImpl {
	return &ConfigOptionImpl{
		f: f,
	}
}

type ConfigOptionImpl struct {
	f func(*configOption)
}

func (c *ConfigOptionImpl) apply() func(db *gorm.DB) *gorm.DB {
	option := &configOption{}
	c.f(option)
	switch option.Column {
	case EnvColumn:
		return func(db *gorm.DB) *gorm.DB {
			if option.Value == OptionAll {
				return db.Preload("Envs")
			} else {
				return db.Preload("Envs", "envs.name = ?", option.Value)
			}
		}
	case ItemColumn:
		return func(db *gorm.DB) *gorm.DB {
			if option.Value == OptionAll {
				return db.Preload("Items")
			} else {
				return db.Preload("Items", "items.name = ?", option.Value)
			}
		}
	default:
		return nil
	}
}
