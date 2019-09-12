package global

import (
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/libtools/clients/client_id"
	"github.com/johnnyeven/libtools/config_agent"
	"github.com/johnnyeven/libtools/courier/client"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/courier/transport_teleport"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/libtools/sqlx/mysql"
	"github.com/johnnyeven/service-vehicle-robot/database"
	"github.com/johnnyeven/service-vehicle-robot/modules"
	"github.com/johnnyeven/service-vehicle-robot/modules/middlewares"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
)

func init() {
	servicex.SetServiceName("service-vehicle-robot")
	servicex.ConfP(&Config)

	database.DBRobot.MustMigrateTo(Config.MasterDB.Get(), !servicex.AutoMigrate)
}

var Config = struct {
	Log *log.Log

	ServeTeleport *transport_teleport.ServeTeleport
	ServeHTTP     transport_http.ServeHTTP

	MasterDB *mysql.MySQL
	SlaveDB  *mysql.MySQL

	ClientID    *client_id.ClientID
	ConfigAgent *config_agent.Agent
	NodeManager *modules.NodeManager `conf:"-"`

	COCOModel *models.COCOObjectDetectiveModel

	RobotConfiguration RobotConfiguration
}{
	Log: &log.Log{
		Level: "DEBUG",
	},

	ServeTeleport: &transport_teleport.ServeTeleport{
		Port: 9090,
		Plugins: []tp.Plugin{
			&middlewares.AuthPlugin{},
		},
	},
	ServeHTTP: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},

	MasterDB: &mysql.MySQL{
		Name:     "robot",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Host:     "localhost",
	},
	SlaveDB: &mysql.MySQL{
		Name:     "robot-readonly",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Host:     "localhost",
	},

	ClientID: &client_id.ClientID{
		Client: client.Client{
			Host: "service-id.profzone.service.profzone.net",
		},
	},
	ConfigAgent: &config_agent.Agent{
		Host:               "service-configurations.profzone.service.profzone.net",
		PullConfigInterval: 60,
		StackID:            124,
	},
	NodeManager: &modules.NodeManager{},

	COCOModel: &models.COCOObjectDetectiveModel{
		ModelPath: "./config/mobilenet",
	},

	RobotConfiguration: RobotConfiguration{},
}
