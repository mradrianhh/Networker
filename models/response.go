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

// MessageType returns the messagetype of the response.
func (response Response) MessageType() networker.MessageType {
	return response.messageType
}

// ResponseCode returns the response code of the response.
func (response Response) ResponseCode() networker.ResponseCode {
	return response.responseCode
}
