package authorization

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/errors"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (a *Authorization) Auth(req *models.AuthRequest) (token []byte, status *tp.Status) {
	node, err := global.Config.NodeManager.GetNodeByKey(req.Key)
	if err != nil {
		return nil, tp.NewStatus(int32(errors.Forbidden), "", errors.Forbidden.StatusError())
	}
	node.Session = a.Session()
	if node.NodeType == types.NODE_TYPE__HOST {
		global.Config.NodeManager.SetHostNode(node)
	}
	return
}
