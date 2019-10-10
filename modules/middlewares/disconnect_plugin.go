package middlewares

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
	"github.com/sirupsen/logrus"
)

type DisconnectPlugin struct {
	Mgr *models.NodeManager
}

func (*DisconnectPlugin) Name() string {
	return "DisconnectPlugin"
}

func (p *DisconnectPlugin) PostDisconnect(sess tp.BaseSession) *tp.Status {
	node, _ := p.Mgr.GetNodeBySessionID(sess.ID())
	if node != nil {
		node.IsOnline = false
		node.Session = nil
		if node.NodeType == types.NODE_TYPE__HOST {
			p.Mgr.SetHostNode(nil)
		}
	}

	logrus.Infof("[%s] disconnect", sess.RemoteAddr().String())
	return nil
}
