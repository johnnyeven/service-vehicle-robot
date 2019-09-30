package camera

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (r *Camera) Holder(req *models.CameraHolderRequest) (bool, *tp.Status) {
	if req == nil {
		return false, nil
	}
	node, err := modules.Manager.GetNodeByKey(req.Target)
	if err != nil {
		return false, tp.NewStatus(int32(errors.NotFound), err.Error(), errors.NotFound.StatusError())
	}

	if !node.IsOnline || node.Session == nil || !node.Session.Health() {
		return false, tp.NewStatus(int32(errors.Forbidden), "远程端无法触达", errors.Forbidden.StatusError())
	}

	return true, node.Session.Push("/camera/holder", req)
}
