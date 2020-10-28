package main

import (
	"fmt"

	"github.com/mradrianhh/Networker/pkg/network"
	"github.com/mradrianhh/Networker/pkg/network/models"
)

func main() {
	client := network.NewClient("tcp", "0.0.0.0:1200")
	response, err := client.Request(models.NewRequest(models.USERNAME))
	fmt.Println(response)
	fmt.Println(err)
}
