package models

import "net"

// Handler ...
type Handler func(Message, net.Conn) error
