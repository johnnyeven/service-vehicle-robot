package models

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/service-vehicle-robot/constants/types"
)

type BroadcastRequest struct {
	Port uint16 `json:"port"`
}

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
	Nodes []*Node `json:"nodes"`
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
