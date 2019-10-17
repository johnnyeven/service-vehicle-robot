package main

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
	"github.com/johnnyeven/service-vehicle-robot/operators"
	"github.com/johnnyeven/service-vehicle-robot/routes"
)

func main() {
	servicex.Execute()

	global.Config.ConfigAgent.BindConf(&global.Config.RobotConfiguration)
	global.Config.ConfigAgent.BindBus(global.Config.MessageBus)
	go global.Config.ConfigAgent.Start()

	models.Manager.Init(global.Config.MasterDB.Get())

	routes.InitRouters()
	go global.Config.ServeHTTP.Serve(operators.RootRouter)
	go global.Config.BroadcastManager.Start(global.Config.ServeTeleport.Port)
	defer global.Config.BroadcastManager.Stop()

	defer tp.FlushLogger()
	go tp.GraceSignal()
	tp.SetLoggerLevel2(tp.ERROR)()
	global.Config.ServeTeleport.Start()
}
