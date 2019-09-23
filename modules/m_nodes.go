package modules

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/database"
	"github.com/sirupsen/logrus"
)

var Manager = &NodeManager{}

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

type Node struct {
	// key
	Key string `json:"key"`
	// secret
	Secret string `json:"secret"`
	// 描述
	Comment string `json:"comment"`
	// 端类型
	NodeType types.NodeType `json:"nodeType"`
	// peer
	Session tp.CtxSession `json:"-"`
	// Token
	Token string `json:"token"`
	// 是否在线
	IsOnline bool `json:"isOnline"`
}

func (n *Node) GenerateToken() string {
	data := n.Key + uuid.New().String()
	n.Token = base64.StdEncoding.EncodeToString([]byte(data))
	return n.Token
}

type NodeManager struct {
	hostNode    *Node
	nodesKeyMap map[string]*Node
	db          *sqlx.DB
}

func (mgr *NodeManager) Init(db *sqlx.DB) {
	mgr.db = db
	model := &database.Nodes{}
	list, _, err := model.FetchList(db, -1, 0)
	if err != nil {
		logrus.Panicf("[NodeManager] Init err: %v", err)
	}

	mgr.nodesKeyMap = make(map[string]*Node)
	for _, node := range list {
		item := Node{
			Key:      node.Key,
			Secret:   node.Secret,
			Comment:  node.Comment,
			NodeType: node.NodeType,
		}
		mgr.nodesKeyMap[node.Key] = &item
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

	item := Node{
		Key:      model.Key,
		Secret:   model.Secret,
		Comment:  model.Comment,
		NodeType: model.NodeType,
	}
	mgr.nodesKeyMap[model.Key] = &item
	return nil
}

func (mgr *NodeManager) GetNodeBySessionID(id string) (*Node, error) {
	for _, node := range mgr.nodesKeyMap {
		if node.Session != nil && id == node.Session.ID() {
			return node, nil
		}
	}

	return nil, nil
}

func (mgr *NodeManager) GetNodeByKey(key string) (node *Node, err error) {
	if v, ok := mgr.nodesKeyMap[key]; ok {
		return v, nil
	}
	model := &database.Nodes{
		Key: key,
	}
	err = model.FetchByKey(mgr.db)
	if err != nil {
		return nil, err
	}
	node = &Node{
		Key:      model.Key,
		Secret:   model.Secret,
		Comment:  model.Comment,
		NodeType: model.NodeType,
	}
	mgr.nodesKeyMap[model.Key] = node
	return
}

func (mgr *NodeManager) GetRobotNode() (nodes []*Node) {
	for _, n := range mgr.nodesKeyMap {
		if n.NodeType == types.NODE_TYPE__ROBOT {
			nodes = append(nodes, n)
		}
	}
	return
}

func (mgr *NodeManager) SetHostNode(node *Node) {
	mgr.hostNode = node
}

func (mgr *NodeManager) GetHostNode() *Node {
	return mgr.hostNode
}
