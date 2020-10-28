package models

// Response is the type. Use NewResponse to initialize a new response-object.
type Response struct {
	MessageType  MessageType
	ResponseCode ResponseCode
}

// NewResponse returns a new response-object with the given responsecode.
func NewResponse(responseCode ResponseCode) Response {
	return Response{
		MessageType:  RESPONSE,
		ResponseCode: responseCode,
	}
}

// GetMessageType ...
func (response Response) GetMessageType() MessageType {
	return response.MessageType
}
