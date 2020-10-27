package models

import (
	networker "github.com/mradrianhh/Networker"
)

// Message represents the object being passed between server and client.
type Message interface {
	MessageType() networker.MessageType
}
