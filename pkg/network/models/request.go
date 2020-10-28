package models

// Request ...
type Request struct {
	messageType MessageType
	requestCode RequestCode
}

// NewRequest returns a new request-object with the given requestcode.
func NewRequest(requestCode RequestCode) Request {
	return Request{
		messageType: RESPONSE,
		requestCode: requestCode,
	}
}

// MessageType returns the messagetype of the request.
func (request Request) MessageType() MessageType {
	return request.messageType
}

// RequestCode returns the requestcode of the request.
func (request Request) RequestCode() RequestCode {
	return request.requestCode
}
