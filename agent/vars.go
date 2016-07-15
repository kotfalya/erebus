package agent

import "flag"

const (
	AGENT_CHECK_INTERVAL = "1s"
	AGENT_CHECK_TIMEOUT  = "100ms"
	AGENT_CHECK_PORT     = 14142
	AGENT_CONN_TYPE      = "tcp"
	AGENT_PORT           = 14141
)

var (
	agentIP = flag.String("agentIP", "127.0.0.1", "agent ip address")
)
