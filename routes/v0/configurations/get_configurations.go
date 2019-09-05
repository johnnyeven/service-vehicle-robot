package configurations

import (
	"context"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/httplib"
	"github.com/johnnyeven/service-vehicle-robot/database"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetConfigurations{}))
}

// 获取配置
type GetConfigurations struct {
	httpx.MethodGet
	// StackID
	StackID uint64 `name:"stackID,string" in:"query"`
	httplib.Pager
}

func (req GetConfigurations) Path() string {
	return ""
}

func (req GetConfigurations) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	resp, count, err := modules.FetchConfiguration(req.StackID, req.Size, req.Offset, db)
	if err != nil {
		logrus.Errorf("[GetConfigurations] modules.FetchConfiguration err: %v, req: %+v", err, req)
		return
	}
	return GetConfigurationResult{
		Data:  resp,
		Total: count,
	}, nil
}

type GetConfigurationResult struct {
	Data  database.ConfigurationList `json:"data"`
	Total int32                      `json:"total"`
}
