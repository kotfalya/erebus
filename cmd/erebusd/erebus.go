package main

import "github.com/kotfalya/erebus/bootstrap"

func main() {
	server := bootstrap.NewServer()
	server.Stop()
}
