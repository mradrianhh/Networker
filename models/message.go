package models

import (
	"fmt"
	"strings"

	networker "github.com/mradrianhh/Networker"
)

// Message represents the object being passed between server and client.
type Message struct {
	MessageType  networker.MessageType
	Message      string
	ResponseCode networker.ResponseCode
	Token        string
}

// NilMessage is an empty message object.
var NilMessage = Message{}

// NewResponseMessage returns a new message-instance for sending a response.
func NewResponseMessage(messageType networker.MessageType, message string, responseCode networker.ResponseCode, token string) Message {
	return Message{MessageType: messageType, Message: message, ResponseCode: responseCode, Token: token}
}

// NewMessage returns a new message-instance for sending a general message.
func NewMessage(messageType networker.MessageType, message string) Message {
	return Message{MessageType: messageType, Message: message}
}

// NewAuthenticatedMessage returns a new message-instance with a token for sending an authenticated message.
func NewAuthenticatedMessage(messageType networker.MessageType, message string, token string) Message {
	return Message{MessageType: messageType, Message: message, Token: token}
}

// Present prints the message to the screen.
func (message Message) Present(index int) {
	fmt.Printf("Message %v\n", index)
	fmt.Printf("\tMessage-type: %s\n", string(message.MessageType))
	fmt.Printf("\tMessage: %s\n", strings.ReplaceAll(message.Message, "\n", " "))
	fmt.Printf("\tResponse-code: %s\n", string(message.ResponseCode))
	fmt.Printf("\tToken: %s\n\n", message.Token)
}
