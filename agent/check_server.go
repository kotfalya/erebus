package agent

import (
	"github.com/golang/glog"
	"net"
	"os"
	"strconv"
)

func AgentCheckAddress() string {
	return *agentIP + ":" + strconv.Itoa(AGENT_CHECK_PORT)
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
		
			conn.Close()
		}
	}()
}
