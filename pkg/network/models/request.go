package models

// Request ...
type Request struct {
	MessageType MessageType
	RequestCode RequestCode
}

// NewRequest returns a new request-object with the given requestcode.
func NewRequest(requestCode RequestCode) Request {
	return Request{
		MessageType: REQUEST,
		RequestCode: requestCode,
	}
}

// GetMessageType ...
func (request Request) GetMessageType() MessageType {
	return request.MessageType
}
