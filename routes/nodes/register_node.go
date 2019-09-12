package nodes

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/helper"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(RegisterNode{}))
}

// 注册端
type RegisterNode struct {
	httpx.MethodPost
	Body modules.RegisterNodeBody `name:"body" in:"body"`
}

func (req RegisterNode) Path() string {
	return ""
}

func (req RegisterNode) Output(ctx context.Context) (result interface{}, err error) {
	id, err := helper.NewUniqueID(global.Config.ClientID)
	if err != nil {
		logrus.Errorf("[RegisterNode] helper.NewUniqueID err: %v, request: %+v", err, req.Body)
		return
	}

	err = global.Config.NodeManager.RegisterNode(id, req.Body)
	if err != nil {
		logrus.Errorf("[RegisterNode] NodeManager.RegisterNode err: %v, request: %+v", err, req.Body)
	}
	return
}
