package models

// Response is the type. Use NewResponse to initialize a new response-object.
type Response struct {
	messageType  MessageType
	responseCode ResponseCode
}

// NewResponse returns a new response-object with the given responsecode.
func NewResponse(responseCode ResponseCode) Response {
	return Response{
		messageType:  RESPONSE,
		responseCode: responseCode,
	}
}

// MessageType returns the messagetype of the response.
func (response Response) MessageType() MessageType {
	return response.messageType
}

// ResponseCode returns the response code of the response.
func (response Response) ResponseCode() ResponseCode {
	return response.responseCode
}
