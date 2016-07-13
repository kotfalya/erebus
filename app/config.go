package app

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/kotfalya/erebus/driver"
	"strings"
)

type Config struct {
	NodeName    string
	ServiceName string
	GroupName   string
	consul      *api.Client
}

func NewConsulConfig() *Config {
	flag.Parse()
	if *serviceName == "" {
		panic("serviceName is requered argument")
	}

	config := &Config{
		*nodeName,
		*serviceName,
		*groupName,
		driver.NewConsulClientNotPooled(),
	}

	if err := config.load(); err != nil {
		panic(err)
	}

	return config
}

func (c *Config) FullName() (fullName string) {
	fullName = c.ServiceName

	if c.GroupName != "" {
		fullName += "-" + c.GroupName
	}

	return
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
