package main

import "github.com/elgong/elgongRPC/server"

func main() {
	server := server.NewRPCServer()

	server.Server()
}
