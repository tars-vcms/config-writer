package configrule

import (
	"gorm.io/gorm"
	"time"
)

type ConfigRule struct {
	ID           uint64
	Key          string
	Value        string
	ConfigItemID uint64 `gorm:"index"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
