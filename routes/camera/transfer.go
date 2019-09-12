package camera

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (r *Camera) Transfer(req *models.CameraRequest) *tp.Status {
	mgr := global.Config.NodeManager
	host := mgr.GetHostNode()
	if host == nil {
		return nil
	}
	stat := host.Session.Push("/camera/transfer", req)
	if !stat.OK() {
		return stat
	}
	return nil
}
