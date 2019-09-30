package camera

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (r *Camera) Transfer(req *models.CameraRequest) (bool, *tp.Status) {
	mgr := modules.Manager
	host := mgr.GetHostNode()
	if host == nil {
		return true, nil
	}
	stat := host.Session.Push("/camera/transfer", req)
	if !stat.OK() {
		return false, stat
	}
	return true, nil
}
