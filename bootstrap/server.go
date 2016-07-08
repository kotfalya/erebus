package bootstrap

import "github.com/golang/glog"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	glog.Infoln("hi")
}

func (s *Server) Stop() {
	glog.Infoln("stop")
}
