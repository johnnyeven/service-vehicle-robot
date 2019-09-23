package nodes

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func (r *Nodes) Robot(req *models.GetNodesRequest) (models.NodesResponse, *tp.Status) {
	nodes := modules.Manager.GetRobotNode()
	return models.NodesResponse{
		Nodes: nodes,
	}, nil
}
