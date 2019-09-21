package models

import "github.com/johnnyeven/service-vehicle-robot/modules"

type AuthRequestHeader struct {
	Token string `json:"token"`
}

type AuthRequest struct {
	Key string `json:"key"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type CameraRequest struct {
	AuthRequestHeader
	Frame []byte `json:"frame"`
}

type GetNodesRequest struct {
	AuthRequestHeader
}

type NodesResponse struct {
	Nodes []*modules.Node `json:"nodes"`
}
