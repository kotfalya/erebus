package main

import (
	"github.com/kotfalya/erebus/agent"
	"github.com/kotfalya/erebus/app"
	"time"
)

func main() {
	cfg := app.NewConsulConfig()
	// check if agent registered on consul
	agent.CheckAndInit(cfg)

	// consul tcp check
	agent.AgentCheckTcpStart()
	
	for {
		select {
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
