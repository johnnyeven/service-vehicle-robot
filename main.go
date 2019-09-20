package main

import (
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/operators"
	"github.com/johnnyeven/service-vehicle-robot/routes"
)

func main() {
	servicex.Execute()

	global.Config.ConfigAgent.BindConf(&global.Config.RobotConfiguration)
	global.Config.ConfigAgent.Start()

	global.Config.NodeManager.Init(global.Config.MasterDB.Get())

	routes.InitRouters()
	go global.Config.ServeHTTP.Serve(operators.RootRouter)
	global.Config.ServeTeleport.Start()
}
