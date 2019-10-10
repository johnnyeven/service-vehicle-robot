package nodes

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func init() {
	Router.Register(courier.NewRouter(GetNodes{}))
}

// 获取节点
type GetNodes struct {
	httpx.MethodGet
}

func (req GetNodes) Path() string {
	return ""
}

func (req GetNodes) Output(ctx context.Context) (result interface{}, err error) {
	result = models.Manager.GetRobotNode()
	return
}
