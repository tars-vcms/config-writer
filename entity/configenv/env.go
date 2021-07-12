package configenv

import (
	"github.com/tars-vcms/config-writer/entity/configitem"
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
	"gorm.io/gorm"
	"time"
)

type ConfigEnv struct {
	ID        uint64
	Items     []configitem.ConfigItem
	ConfigID  uint64 `gorm:"index"`
	Name      string `gorm:"uniqueIndex:env_index;size:15"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *ConfigEnv) ToPB() *config_writer.ConfigEnv {
	return &config_writer.ConfigEnv{
		ID:        c.ID,
		ConfigID:  c.ConfigID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt.Unix(),
	}
}
