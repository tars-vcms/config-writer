package config

import (
	"github.com/tars-vcms/config-writer/entity/configenv"
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
	"gorm.io/gorm"
	"time"
)

type Status uint16

const (
	Online  Status = 1
	Offline Status = 2
)

var StatusPBMap = map[Status]config_writer.STATUS_TYPE{
	Online:  config_writer.STATUS_TYPE_ONLINE,
	Offline: config_writer.STATUS_TYPE_OFFLINE,
}

var PBStatusMap = map[config_writer.STATUS_TYPE]Status{
	config_writer.STATUS_TYPE_ONLINE:  Online,
	config_writer.STATUS_TYPE_OFFLINE: Offline,
}

type Config struct {
	ID        uint64
	Name      string `gorm:"uniqueIndex;size:15"`
	Secret    string
	Status    Status
	Comment   string
	Envs      []configenv.ConfigEnv
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *Config) ToPB() *config_writer.Config {
	return &config_writer.Config{
		ID:        c.ID,
		Secret:    c.Secret,
		Status:    StatusPBMap[c.Status],
		Name:      c.Name,
		Comment:   c.Comment,
		Env:       c.EnvsToPB(),
		CreatedAt: c.CreatedAt.Unix(),
	}
}

func (c *Config) EnvsToPB() []*config_writer.ConfigEnv {
	var envs []*config_writer.ConfigEnv
	for _, value := range c.Envs {
		envs = append(envs, value.ToPB())
	}
	return envs
}
