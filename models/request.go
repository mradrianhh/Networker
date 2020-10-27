package models

import networker "github.com/mradrianhh/Networker"

// Request ...
type Request struct {
	messageType networker.MessageType
	requestCode networker.RequestCode
}

// NewRequest returns a new request-object with the given requestcode.
func NewRequest(requestCode networker.RequestCode) Request {
	return Request{
		messageType: networker.RESPONSE,
		requestCode: requestCode,
	}
}

// MessageType returns the messagetype of the request.
func (request Request) MessageType() networker.MessageType {
	return request.messageType
}

// RequestCode returns the requestcode of the request.
func (request Request) RequestCode() networker.RequestCode {
	return request.requestCode
}
