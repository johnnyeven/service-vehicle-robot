package authorization

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (a *Authorization) Auth(req *models.AuthRequest) (resp models.AuthResponse, status *tp.Status) {
	node, err := modules.Manager.GetNodeByKey(req.Key)
	if err != nil {
		status = tp.NewStatus(int32(errors.Forbidden), "", errors.Forbidden.StatusError())
		return
	}
	node.Session = a.Session()
	if node.NodeType == types.NODE_TYPE__HOST {
		modules.Manager.SetHostNode(node)
	}
	node.IsOnline = true

	resp.Token = node.GenerateToken()
	return
}
