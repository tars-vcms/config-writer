package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/tars-vcms/config-writer/entity/config"
	"github.com/tars-vcms/config-writer/entity/configitem"
	"github.com/tars-vcms/config-writer/entity/errcode"
	"github.com/tars-vcms/config-writer/repo/cfg"
	"github.com/tars-vcms/config-writer/repo/remotecfg"
	"github.com/tars-vcms/vcms-common/errs"
	config_writer "github.com/tars-vcms/vcms-protocol/config/config-writer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigServerImpl struct {
	db   *gorm.DB
	rcfg remotecfg.RemoteConfig
	cfg  cfg.Config
}

func newConfigServerImpl() *ConfigServerImpl {
	rcfg := remotecfg.New()
	dsn, err := rcfg.GetDatabaseDSN()
	if err != nil {
		panic("[database]dsn:" + err.Error())
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[database]dsn:" + err.Error())
	}
	return &ConfigServerImpl{
		db:   db,
		rcfg: rcfg,
		cfg:  cfg.New(db),
	}
}

func (c *ConfigServerImpl) CreateConfig(ctx context.Context, name string) (*config_writer.Config, error) {

	if err := c.checkCreateConfigParam(name); err != nil {
		return nil, err
	}

	cfgEntity := &config.Config{
		Name:   name,
		Status: config.Offline,
		Secret: c.genSecret(),
	}

	if err := c.cfg.InputConfig(ctx, cfgEntity); err != nil {
		return nil, err
	}

	return &config_writer.Config{
		ID:        cfgEntity.ID,
		Secret:    cfgEntity.Secret,
		Status:    config.StatusPBMap[cfgEntity.Status],
		Comment:   cfgEntity.Comment,
		CreatedAt: cfgEntity.CreatedAt.Unix(),
	}, nil
}

func (c *ConfigServerImpl) UpdateConfig(ctx context.Context, pbConfig *config_writer.Config) error {
	if err := c.checkUpdateConfigParam(pbConfig); err != nil {
		return err
	}
	cfgEntities, err := c.cfg.GetByIDs([]uint64{pbConfig.ID}, cfg.WithEnv(cfg.OptionAll))
	if err != nil {
		return err
	}
	if len(*cfgEntities) == 0 {
		return errs.Newf(errcode.RetIDErr, "config id not exists")
	}
	cfgEntity := (*cfgEntities)[0]
	//修改部分可修改字段
	cfgEntity.Comment = pbConfig.Comment
	cfgEntity.Status = config.PBStatusMap[pbConfig.Status]
	cfgEntity.Name = "name"
	if err := c.cfg.SaveConfig(ctx, &cfgEntity); err != nil {
		return err
	}
	return nil
}

func (c *ConfigServerImpl) GetConfigs(CIds []uint64, Names []string, withEnv bool) ([]*config_writer.Config, error) {
	CidConfigEntities, err := c.cfg.GetByIDs(CIds, c.convertOption(cfg.WithEnv, cfg.OptionAll, withEnv))
	if err != nil {
		return nil, err
	}
	NameConfigEntities, err := c.cfg.GetByNames(Names, c.convertOption(cfg.WithEnv, cfg.OptionAll, withEnv))
	if err != nil {
		return nil, err
	}
	var configs []*config_writer.Config
	for _, entity := range *CidConfigEntities {
		configs = append(configs, entity.ToPB())
	}
	for _, entity := range *NameConfigEntities {
		configs = append(configs, entity.ToPB())
	}
	return configs, nil
}

func (c *ConfigServerImpl) GetItems() {
	panic("implement me")
}

func (c *ConfigServerImpl) convertOption(f func(name string) cfg.ConfigOption, name string, flag bool) cfg.ConfigOption {
	if flag {
		return f(name)
	}
	return nil
}

func (c *ConfigServerImpl) checkCreateConfigParam(name string) error {
	identLen := len(name)
	if identLen <= 5 || identLen >= 15 {
		return errs.Newf(errcode.RetNameErr, "config ident illegal")
	}
	cfgEntity, err := c.cfg.GetByNames([]string{name})
	if err != nil {
		return err
	}
	if len(*cfgEntity) != 0 {
		return errs.Newf(errcode.RetNameErr, "config name has used")
	}
	return nil
}

func (c *ConfigServerImpl) checkUpdateConfigParam(config *config_writer.Config) error {
	if config.ID == 0 {
		return errs.Newf(errcode.RetIDErr, "config id illegal")
	}
	return nil
}

func (c *ConfigServerImpl) genSecret() string {
	return uuid.NewString()
}

func (c *ConfigServerImpl) convertToPBContents(contentEntity *configitem.ConfigContent) *config_writer.ConfigContent {
	return nil
}
