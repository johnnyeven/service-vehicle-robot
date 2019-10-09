package modules

import (
	"encoding/json"
	"github.com/johnnyeven/service-vehicle-robot/global"
	"github.com/johnnyeven/service-vehicle-robot/modules/models"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type BroadcastManager struct {
	conn net.Conn
	quit chan struct{}
}

func (mgr *BroadcastManager) Init() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 9091,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		logrus.Panic(err)
	}
	mgr.conn = conn
	mgr.quit = make(chan struct{})
}

func (mgr *BroadcastManager) Start() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	broadcast := models.BroadcastRequest{Port: global.Config.ServeTeleport.Port}
	bytes, err := json.Marshal(broadcast)
	if err != nil {
		logrus.Panic(err)
	}

Run:
	for {
		select {
		case <-ticker.C:
			_, err := mgr.conn.Write(bytes)
			if err != nil {
				logrus.Warningf("[BroadcastManager] conn.Write err: %v", err)
				continue
			}
			logrus.Info("[BroadcastManager] send broadcast")
		case <-mgr.quit:
			break Run
		}
	}
}

func (mgr *BroadcastManager) Stop() error {
	logrus.Info("[BroadcastManager] stop broadcast")
	mgr.quit <- struct{}{}
	return mgr.conn.Close()
}
