package main

import (
	"fmt"
	"github.com/tars-vcms/vcms-protocol/config/config-writer"
	"os"

	"github.com/TarsCloud/TarsGo/tars"
)

func main() {
	// Get server cfg
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(ConfigWriterImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("configImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}

	//tars.RegisterServerFilter(ErrorFilter)

	// New servant
	app := new(config_writer.ConfigWriter)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".configObj")

	// Run application
	tars.Run()
}
