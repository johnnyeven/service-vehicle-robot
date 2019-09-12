package modules

import (
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/database"
	"github.com/sirupsen/logrus"
)

type RegisterNodeBody struct {
	// key
	Key string `json:"key"`
	// secret
	Secret string `json:"secret"`
	// 描述
	Comment string `json:"comment"`
	// 端类型
	NodeType types.NodeType `json:"nodeType"`
}

type NodeManager struct {
	nodes    database.NodesList
	nodesMap map[string]database.Nodes
	db       *sqlx.DB
}

func (mgr *NodeManager) Init(db *sqlx.DB) {
	mgr.db = db
	model := &database.Nodes{}
	list, _, err := model.FetchList(db, -1, 0)
	if err != nil {
		logrus.Panicf("[NodeManager] Init err: %v", err)
	}
	mgr.nodes = append(mgr.nodes, list...)
	mgr.nodesMap = make(map[string]database.Nodes)
	for _, node := range mgr.nodes {
		mgr.nodesMap[node.Key] = node
	}
}

func (mgr *NodeManager) RegisterNode(id uint64, body RegisterNodeBody) error {
	model := &database.Nodes{
		NodesID:  id,
		Key:      body.Key,
		Secret:   body.Secret,
		Comment:  body.Comment,
		NodeType: body.NodeType,
	}
	err := model.Create(mgr.db)
	if err != nil {
		return err
	}

	err = model.FetchByNodesID(mgr.db)
	if err != nil {
		return err
	}

	mgr.nodes = append(mgr.nodes, *model)
	mgr.nodesMap[model.Key] = *model
	return nil
}

func (mgr *NodeManager) GetNodeByKey(key string) (model *database.Nodes, err error) {
	if v, ok := mgr.nodesMap[key]; ok {
		return &v, nil
	}
	model = &database.Nodes{
		Key: key,
	}
	err = model.FetchByKey(mgr.db)
	if err != nil {
		return nil, err
	}
	mgr.nodes = append(mgr.nodes, *model)
	mgr.nodesMap[model.Key] = *model
	return
}
