package configitem

import (
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
	"gorm.io/gorm"
	"time"
)

type ConfigContent struct {
	ID           uint64
	Version      uint32 `gorm:"index:item_index"`
	Content      string
	ConfigItemID uint64 `gorm:"index:item_index"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (c *ConfigContent) ToPB() *config_writer.ConfigContent {
	return &config_writer.ConfigContent{
		ID:        c.ID,
		Version:   c.Version,
		Content:   c.Content,
		CreatedAt: c.CreatedAt.Unix(),
	}
}
