package network

import (
	"net"

	"github.com/mradrianhh/Networker/pkg/network/models"
)

// Handler ...
type Handler func(models.Request, net.Conn) error
