package app

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/hashicorp/consul/api"
	"github.com/kotfalya/erebus/driver"
)

type DistributedLock struct {
	name     string
	config   *Config
	consul   *api.Client
	stopLock chan struct{}
}

func NewDistributedLock(name string, config *Config, stopLock chan struct{}) *DistributedLock {
	return &DistributedLock{
		name,
		config,
		driver.NewConsulClient(),
		stopLock,
	}
}

func (dl *DistributedLock) Lock() (leaderCh <-chan struct{}, err error) {
	lock, err := dl.consul.LockOpts(&api.LockOptions{
		Key:         dl.getLockKey(),
		SessionName: dl.getSessionName(),
	})
	if err != nil {
		panic(err)
	}

TRY_LOCK:
	leaderCh, err = lock.Lock(dl.stopLock)
	if err != nil {
		panic(err)
	}

	if leaderCh == nil {
		goto TRY_LOCK
	}

	glog.V(1).Infof("New lock leader, Service %s, Node %s, name %s \n", dl.config.ServiceName, dl.config.NodeName, dl.name)

	return
}

func (dl *DistributedLock) GetLock(name string) (*api.Lock, error) {
	return dl.consul.LockOpts(&api.LockOptions{
		Key:         dl.getLockKey(),
		SessionName: dl.getSessionName(),
	})
}

func (dl *DistributedLock) getSessionName() string {
	if dl.config.GroupName != "" {
		return fmt.Sprintf("session-%s-%s-%s", dl.config.GroupName, dl.config.ServiceName, dl.name)
	} else {
		return fmt.Sprintf("session-%s-%s", dl.config.ServiceName, dl.name)
	}
}

func (dl *DistributedLock) getLockKey() string {
	return fmt.Sprintf("service/default/%s/lock/%s", dl.config.ServiceName, dl.name)
}
