package main

import (
	"github.com/mradrianhh/Networker/pkg/network"
)

func main() {
	server := network.NewServer("tcp", "0.0.0.0:1200")
	server.Listen()
}
