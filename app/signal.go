package app

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/kotfalya/erebus/driver"
	"sync/atomic"
)

type Signal struct {
	config    *Config
	consul    *api.Client
	Signal    chan string
	err       chan error
	stop      chan bool
	waitIndex uint64
}

func NewSignal(config *Config) *Signal {
	return &Signal{
		config,
		driver.NewConsulClient(),
		make(chan string),
		make(chan error),
		make(chan bool),
		uint64(0),
	}
}

func (s *Signal) Start() {
	if _, err := s.loadSignal(); err != nil {
		panic(err)
	}

	go s.listenSignal()
}

func (s *Signal) Stop() {
	close(s.stop)
}

func (s *Signal) loadSignal() (string, error) {
	pair, _, err := s.consul.KV().Get(s.getSignalKey(), &api.QueryOptions{WaitIndex: s.waitIndex})
	if err != nil {
		return "", err
	}

	if pair != nil {
		if s.waitIndex == pair.ModifyIndex {
			return "", nil
		}
		atomic.StoreUint64(&s.waitIndex, pair.ModifyIndex)

		return string(pair.Value), nil
	} else {
		_, err := s.initSignal()
		if err != nil {
			return "", err
		}

		return "", nil
	}

}

func (s *Signal) initSignal() (*api.KVPair, error) {
	key := &api.KVPair{Value: []byte("init"), Key: s.getSignalKey()}
	_, err := s.consul.KV().Put(key, nil)

	return key, err
}

func (s *Signal) listenSignal() {
	for {
		select {
		case <-s.stop:
			return
		default:
			if signal, err := s.loadSignal(); err != nil {
				s.err <- err

				return
			} else if signal != "" {
				s.Signal <- signal
			}
		}
	}
}

func (s *Signal) getSignalKey() string {
	return fmt.Sprintf("service/node/%s/%s/signal", s.config.NodeName, s.config.ServiceName)
}
