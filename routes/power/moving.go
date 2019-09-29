package power

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (p *Power) Moving(req *models.PowerMovingRequest) *tp.Status {
	if req == nil {
		return nil
	}
	node, err := modules.Manager.GetNodeByKey(req.Target)
	if err != nil {
		return tp.NewStatus(int32(errors.NotFound), err.Error(), errors.NotFound.StatusError())
	}

	if !node.IsOnline || node.Session == nil || !node.Session.Health() {
		return tp.NewStatus(int32(errors.Forbidden), "远程端无法触达", errors.Forbidden.StatusError())
	}

	return node.Session.Push("/power/moving", req)
}
