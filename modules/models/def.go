package models

import (
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
	"github.com/johnnyeven/service-vehicle-robot/modules"
)

type AuthRequestHeader struct {
	Token string `json:"token"`
}

type TargetRequestHeader struct {
	Target string `json:"target"`
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

type PowerMovingRequest struct {
	AuthRequestHeader
	TargetRequestHeader
	Direction types.MovingDirection `json:"direction"`
	Speed     float64               `json:"speed"`
}

type CameraHolderRequest struct {
	AuthRequestHeader
	TargetRequestHeader
	Direction types.HolderDirection `json:"direction"`
	Offset    float64               `json:"offset"`
}
