package nodes

import "github.com/johnnyeven/libtools/courier"

var Router = courier.NewRouter(NodeGroup{})

type NodeGroup struct {
	courier.EmptyOperator
}

func (NodeGroup) Path() string {
	return "/nodes"
}
