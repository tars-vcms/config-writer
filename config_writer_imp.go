package main

import (
	"context"
	"github.com/tars-vcms/config-writer/logic"
	"github.com/tars-vcms/vcms-common/errs"
	. "github.com/tars-vcms/vcms-protocol/config/config-writer"
)

type ConfigWriterImp struct {
	server logic.ConfigServer
}

// Init servant init
func (imp *ConfigWriterImp) Init() error {
	//initialize servant here:
	//...
	imp.server = logic.NewConfigServer()
	return nil
}

// Destroy servant destory
func (imp *ConfigWriterImp) Destroy() {
	//destroy servant here:
	//...
}
func (imp *ConfigWriterImp) InsertConfig(
	tarsCtx context.Context,
	input InsertConfigRequest,
) (output InsertConfigReply, err error) {
	config, err := imp.server.CreateConfig(context.Background(), input.GetName())
	if err != nil {
		return output, errs.HandleError(tarsCtx, err)
	}
	output.Config = config
	return output, nil
}

func (imp *ConfigWriterImp) GetConfigs(ctx context.Context, input GetConfigsRequest) (output GetConfigsReply, err error) {
	panic("implement me")
}

func (imp *ConfigWriterImp) UpdateConfig(ctx context.Context, input UpdateConfigRequest) (output UpdateConfigReply, err error) {
	panic("implement me")
}
