package database

import (
	"github.com/johnnyeven/libtools/sqlx/presets"
)

//go:generate libtools gen model Configuration --database DBRobot --table-name t_configuration --with-comments
//go:generate libtools gen tag Configuration --defaults=true
// @def primary ID
// @def unique_index U_configuration_id ConfigurationID
// @def unique_index U_stack_key StackID Key
// @def index I_stack StackID
type Configuration struct {
	presets.PrimaryID
	// 业务ID
	ConfigurationID uint64 `db:"F_configuration_id" json:"configurationID,string" sql:"bigint unsigned NOT NULL DEFAULT '0'"`
	// StackID
	StackID uint64 `db:"F_stack_id" json:"stackID,string" sql:"bigint unsigned NOT NULL DEFAULT '0'"`
	// Key
	Key string `db:"F_key" json:"key" sql:"varchar(255) NOT NULL DEFAULT ''"`
	// Value
	Value string `db:"F_value" json:"value" sql:"varchar(255) NOT NULL DEFAULT ''"`

	presets.OperateTime
	presets.SoftDelete
}
