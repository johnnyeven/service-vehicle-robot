package configurations

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(CreateConfig{}))
}

// 创建配置
type CreateConfig struct {
	httpx.MethodPost
	Body modules.CreateConfigurationBody `name:"body" in:"body"`
}

func (req CreateConfig) Path() string {
	return ""
}

func (req CreateConfig) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.CreateConfiguration(req.Body, db, global.Config.ClientID)
	if err != nil {
		logrus.Errorf("[CreateConfig] modules.CreateConfiguration err: %v, req: %+v", err, req.Body)
	}
	return
}
