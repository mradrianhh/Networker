package models

// Message represents the object being passed between server and client.
type Message interface {
	MessageType() MessageType
}
