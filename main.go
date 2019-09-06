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

	routes.InitRouters()
	global.Config.ServeTeleport.Start()
}
