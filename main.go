package main

import (
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/routes"
)

func main() {
	servicex.Execute()

	global.Config.ConfigAgent.BindConf(&global.Config.RobotConfiguration)
	global.Config.ConfigAgent.Start()

	go global.Config.ServerHTTP.Serve(routes.RootRouter)
	global.Config.ServerGRPC.Serve(routes.RootRouter)
}
