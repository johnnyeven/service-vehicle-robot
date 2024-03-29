package routes

import (
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/routes/authorization"
	"github.com/johnnyeven/service-vehicle-robot/routes/camera"
	"github.com/johnnyeven/service-vehicle-robot/routes/detection"
	"github.com/johnnyeven/service-vehicle-robot/routes/nodes"
	"github.com/johnnyeven/service-vehicle-robot/routes/power"
)

func InitRouters() {
	server := global.Config.ServeTeleport
	server.RegisterCallRouter(&authorization.Authorization{})
	server.RegisterCallRouter(&nodes.Nodes{})
	server.RegisterCallRouter(&detection.Detection{})
	server.RegisterCallRouter(&camera.Camera{})
	server.RegisterPushRouter(&power.Power{})
}
