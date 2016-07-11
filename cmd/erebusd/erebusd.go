package main

import (
	"github.com/kotfalya/erebus/bootstrap"
	"github.com/kotfalya/erebus/app"
)


func main() {
	cfg := app.NewConsulConfig()
	cfg.Parse()

	server := bootstrap.NewServer()
	server.Start()
}
