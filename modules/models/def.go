package models

import "github.com/johnnyeven/service-vehicle-robot/modules"

type AuthRequestHeader struct {
	Token string
}

type CameraRequest struct {
	AuthRequestHeader
	Frame []byte
}

type GetNodesRequest struct {
	AuthRequestHeader
}

type NodesResponse struct {
	Nodes []*modules.Node
}
