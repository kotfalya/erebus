package agent

import (
	"github.com/golang/glog"
	"net"
	"os"
	"strconv"
)

const (
	AGENT_CHECK_INTERVAL = "1s"
	AGENT_CHECK_TIMEOUT  = "100ms"
	AGENT_CONN_TYPE      = "tcp"
)

func AgentCheckAddress() string {
	return *agentIP + ":" + strconv.Itoa(*agentCheckPort)
}

func AgentCheckTcpStart() {
	go func() {
		l, err := net.Listen(AGENT_CONN_TYPE, AgentCheckAddress())
		if err != nil {
			glog.Errorln("Error listening:", err.Error())
			os.Exit(1)
		}

		defer l.Close()
		glog.V(1).Infoln("Listening on " + AgentCheckAddress())
		for {
			conn, err := l.Accept()
			if err != nil {
				glog.Errorln("Error accepting: ", err.Error())
			}

			glog.V(3).Infoln("connected")
			conn.Close()
		}
	}()
}
