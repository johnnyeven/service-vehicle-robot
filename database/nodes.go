package database

import (
	"github.com/johnnyeven/libtools/sqlx/presets"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
)

//go:generate libtools gen model Nodes --database DBRobot --table-name t_nodes --with-comments
//go:generate libtools gen tag Nodes --defaults=true
// @def primary ID
// @def unique_index U_nodes_id NodesID
// @def unique_index U_key Key
type Nodes struct {
	presets.PrimaryID
	// 业务ID
	NodesID uint64 `db:"F_nodes_id" json:"nodesID,string" sql:"bigint(64) unsigned NOT NULL"`
	// key
	Key string `db:"F_key" json:"key" sql:"varchar(255) NOT NULL DEFAULT ''"`
	// secret
	Secret string `db:"F_secret" json:"secret" sql:"varchar(255) NOT NULL DEFAULT ''"`
	// 描述
	Comment string `db:"F_comment" json:"comment" sql:"varchar(255) NOT NULL DEFAULT ''"`
	// 端类型
	NodeType types.NodeType `db:"F_node_type" json:"nodeType" sql:"tinyint unsigned NOT NULL DEFAULT '0'"`

	presets.OperateTime
	presets.SoftDelete
}
