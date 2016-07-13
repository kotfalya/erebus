package agent

import (
	"github.com/golang/glog"
	"github.com/hashicorp/consul/api"
	"github.com/kotfalya/erebus/app"
	"github.com/kotfalya/erebus/driver"
)

func CheckAndInit(config *app.Config) {
	consul := driver.NewConsulClientNotPooled()

	services, err := consul.Agent().Services()
	if err != nil {
		panic(err)
	}
	
	if _, ok := services[config.FullName()]; ok {
		return
	}

	err = consul.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:   config.FullName(),
		Name: config.FullName(),
		Tags: []string{app.APP_LEVEL_SYSTEM},
	})

	if err != nil {
		panic(err)
	}

	err = consul.Agent().CheckRegister(&api.AgentCheckRegistration{
		Name:      config.FullName() + "-tcp-check",
		ServiceID: config.FullName(),
		AgentServiceCheck: api.AgentServiceCheck{
			TCP:      AgentCheckAddress(),
			Interval: AGENT_CHECK_INTERVAL,
			Timeout:  AGENT_CHECK_TIMEOUT,
		},
	})

	if err != nil {
		panic(err)
	}

	glog.V(1).Infoln("Agent was registerd on consul with name:", config.FullName())
}
