package nodes

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (r *Nodes) Robot(req *models.GetNodesRequest) (models.NodesResponse, *tp.Status) {
	nodes := global.Config.NodeManager.GetRobotNode()
	return models.NodesResponse{
		Nodes: nodes,
	}, nil
}
