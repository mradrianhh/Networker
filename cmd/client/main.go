package main

import (
	"github.com/mradrianhh/Networker/pkg/network"
	"github.com/mradrianhh/Networker/pkg/network/models"
)

func main() {
	client := network.NewClient("tcp", "0.0.0.0:1200")
	client.Request(models.NewRequest(models.USERNAME))
}
