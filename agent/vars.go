package agent

import "flag"

var (
	agentIP        = flag.String("agentIP", "127.0.0.1", "agent ip address")
	agentPort      = flag.Int("agentPort", 14141, "agent tcp port")
	agentCheckPort = flag.Int("agentCheckPort", 14142, "agent check tcp port")
)
