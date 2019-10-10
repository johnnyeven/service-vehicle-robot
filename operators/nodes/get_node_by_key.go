package nodes

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetNodeByKey{}))
}

// 根据Key获取节点
type GetNodeByKey struct {
	httpx.MethodGet
	// Key
	Key string `name:"key" in:"path"`
}

func (req GetNodeByKey) Path() string {
	return "/:key"
}

func (req GetNodeByKey) Output(ctx context.Context) (result interface{}, err error) {
	result, err = models.Manager.GetNodeByKey(req.Key)
	if err != nil {
		logrus.Errorf("[GetNodeByKey] NodeManager.GetNodeByKey err: %v, request: %+v", err, req)
	}
	return
}
