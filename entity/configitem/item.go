package configitem

import (
	"github.com/tars-vcms/config-writer/entity/configrule"
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
	"gorm.io/gorm"
	"time"
)

type ItemType uint16

const (
	Text ItemType = 1
	Json ItemType = 2
	Yaml ItemType = 3
)

var ItemPBMap = map[ItemType]config_writer.ITEM_TYPE{
	Text: config_writer.ITEM_TYPE_TEXT,
	Json: config_writer.ITEM_TYPE_JSON,
	Yaml: config_writer.ITEM_TYPE_YAML,
}

type ConfigItem struct {
	ID            uint64 `gorm:"primaryKey"`
	Rules         []configrule.ConfigRule
	ConfigEnvID   uint64 `gorm:"index:config_index"`
	Name          string `gorm:"uniqueIndex:item_index;size:15"`
	ActiveVersion uint64
	Type          ItemType
	Contents      []ConfigContent
	Comment       string
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (c *ConfigItem) ToPB() *config_writer.ConfigItem {
	return &config_writer.ConfigItem{
		ID:            c.ID,
		ConfigEnvID:   c.ConfigEnvID,
		ActiveVersion: c.ActiveVersion,
		Name:          c.Name,
		Type:          ItemPBMap[c.Type],
		CreatedAt:     c.CreatedAt.Unix(),
	}
}

func (c *ConfigItem) ToPBContents() []*config_writer.ConfigContent {
	var envs []*config_writer.ConfigContent
	for _, value := range c.Contents {
		envs = append(envs, value.ToPB())
	}
	return envs
}
