package main

import (
	"config-server/tars-protocol/vcms"
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("vcms.cfg-server.configObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(vcms.Config)
	comm.StringToProxy(obj, app)
	var out, i int32
	i = 123
	ret, err := app.Add(i, i*2, &out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret, out)
}
