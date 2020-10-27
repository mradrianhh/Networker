package models

import networker "github.com/mradrianhh/Networker"

// Response is the type. Use NewResponse to initialize a new response-object.
type Response struct {
	messageType  networker.MessageType
	responseCode networker.ResponseCode
}

// NewResponse returns a new response-object with the given responsecode.
func NewResponse(responseCode networker.ResponseCode) Response {
	return Response{
		messageType:  networker.RESPONSE,
		responseCode: responseCode,
	}
}

// MessageType ...
func (response Response) MessageType() networker.MessageType {
	return response.messageType
}
