package app

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/kotfalya/erebus/driver"
	"strings"
)

var nodeName = flag.String("nodeName", "node_default", "Node name in cluster")
var serviceName = flag.String("serviceName", "", "Name of service")
var groupName = flag.String("groupName", "", "Name of application")

type Config struct {
	NodeName    string
	ServiceName string
	GroupName   string
	consul      *api.Client
}

func NewConsulConfig() *Config {
	return &Config{
		*nodeName,
		*serviceName,
		*groupName,
		driver.NewConsulClientNotPooled(),
	}
}

func (c *Config) Parse() {
	flag.Parse()

	if *serviceName == "" {
		panic("serviceName is requered argument")
	}

	if err := c.load(); err != nil {
		panic(err)
	}
}

func (c *Config) load() error {
	if c.GroupName != "" {
		if err := c.loadPrefix(c.getGroupDefaultConfigKeyPrefix()); err != nil {
			return err
		}

		if err := c.loadPrefix(c.getGroupNodeConfigKeyPrefix()); err != nil {
			return err
		}
	}

	if err := c.loadPrefix(c.getServiceDefaultConfigKeyPrefix()); err != nil {
		return err
	}

	if err := c.loadPrefix(c.getServiceNodeConfigKeyPrefix()); err != nil {
		return err
	}

	return nil
}

func (c *Config) loadPrefix(prefix string) error {
	pairs, _, err := c.consul.KV().List(prefix, nil)
	if err != nil {
		return err
	}

	for _, v := range pairs {
		key := c.parseConfigKey(prefix, v.Key)

		if r := flag.Lookup(key); r != nil {
			if err := flag.Set(key, string(v.Value)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Config) getServiceDefaultConfigKeyPrefix() string {
	return fmt.Sprintf("service/default/%s/config/", c.ServiceName)
}

func (c *Config) getServiceNodeConfigKeyPrefix() string {
	return fmt.Sprintf("service/node/%s/%s/config/", c.NodeName, c.ServiceName)
}

func (c *Config) getGroupDefaultConfigKeyPrefix() string {
	return fmt.Sprintf("service/default/%s/config/", c.GroupName)
}

func (c *Config) getGroupNodeConfigKeyPrefix() string {
	return fmt.Sprintf("service/node/%s/%s/config/", c.NodeName, c.GroupName)
}

func (c *Config) parseConfigKey(prefix string, key string) string {
	return strings.Replace(key, prefix, "", 1)
}
