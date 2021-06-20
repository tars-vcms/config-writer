package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"config-server/tars-protocol/vcms"
)

func main() {
	// Get server cfg
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(configImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("configImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(vcms.Config)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".configObj")

	// Run application
	tars.Run()
}
