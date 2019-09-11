package routes

import (
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/routes/authorization"
	"github.com/johnnyeven/service-vehicle-robot/routes/camera"
	"github.com/johnnyeven/service-vehicle-robot/routes/detection"
)

func InitRouters() {
	server := global.Config.ServeTeleport
	server.RegisterCallRouter(&authorization.Authorization{})
	server.RegisterCallRouter(&detection.Detection{})
	server.RegisterPushRouter(&camera.Camera{})
}
