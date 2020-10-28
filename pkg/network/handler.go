package network

import (
	"net"

	"github.com/mradrianhh/Networker/pkg/network/models"
)

// Handler ...
type Handler func(models.Message, net.Conn) error
