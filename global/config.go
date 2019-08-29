package global

import (
	"github.com/johnnyeven/libtools/courier/transport_grpc"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/libtools/sqlx/mysql"
	"github.com/johnnyeven/service-vehicle-robot/database"
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
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	ServerHTTP: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
	ServerGRPC: transport_grpc.ServeGRPC{
		Port: 9000,
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
}
