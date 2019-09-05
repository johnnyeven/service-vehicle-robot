package global

import (
	"github.com/johnnyeven/libtools/clients/client_id"
	"github.com/johnnyeven/libtools/config_agent"
	"github.com/johnnyeven/libtools/courier/client"
	"github.com/johnnyeven/libtools/courier/transport_grpc"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/libtools/sqlx/mysql"
	"github.com/johnnyeven/service-vehicle-robot/database"
	"github.com/johnnyeven/service-vehicle-robot/modules"
)

func init() {
	servicex.SetServiceName("service-vehicle-robot")
	servicex.ConfP(&Config)

	database.DBRobot.MustMigrateTo(Config.MasterDB.Get(), !servicex.AutoMigrate)
}

var Config = struct {
	Log        *log.Log
	ServerHTTP transport_http.ServeHTTP
	ServerGRPC transport_grpc.ServeGRPC

	MasterDB *mysql.MySQL
	SlaveDB  *mysql.MySQL

	ClientID    *client_id.ClientID
	ConfigAgent *config_agent.Agent

	COCOModel *modules.COCOObjectDetectiveModel

	RobotConfiguration RobotConfiguration
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	ServerHTTP: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
	ServerGRPC: transport_grpc.ServeGRPC{
		Port: 9900,
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
		StackID:            123,
	},

	COCOModel: &modules.COCOObjectDetectiveModel{
		ModelPath: "./config/mobilenet",
	},

	RobotConfiguration: RobotConfiguration{},
}
