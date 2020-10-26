package models

import (
	networker "Networker"
	"fmt"
	"strings"
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

// ErrorMessage is the default message when something goes wrong.
var ErrorMessage = Message{MessageType: networker.Confirmation, Message: "error", ResponseCode: networker.NotAccepted}

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

// NewMsgUnauthenticatedToken sends a response message informing the request-sender that the token is not authenticated.
func NewMsgUnauthenticatedToken(token string) Message {
	return NewResponseMessage(networker.Confirmation, token+" is not registered as an authenticated token.", networker.NotAccepted, token)
}

// Present prints the message to the screen.
func (message Message) Present(index int) {
	fmt.Printf("Message %v\n", index)
	fmt.Printf("\tMessage-type: %s\n", string(message.MessageType))
	fmt.Printf("\tMessage: %s\n", strings.ReplaceAll(message.Message, "\n", " "))
	fmt.Printf("\tResponse-code: %s\n", string(message.ResponseCode))
	fmt.Printf("\tToken: %s\n\n", message.Token)
}
