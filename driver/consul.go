package driver

import (
	"flag"
	"github.com/hashicorp/consul/api"
)

var (
	consulAddr = flag.String("consulAddr", "", "Consul agent address")
)

func NewConsulClient() *api.Client {
	conf := api.DefaultConfig()
	if *consulAddr != "" {
		conf.Address = *consulAddr
	}

	client, err := api.NewClient(conf)
	if err != nil {
		panic(err)
	}

	return client
}

func NewConsulClientNotPooled() *api.Client {
	conf := api.DefaultNonPooledConfig()
	if *consulAddr != "" {
		conf.Address = *consulAddr
	}

	client, err := api.NewClient(conf)
	if err != nil {
		panic(err)
	}

	return client
}
