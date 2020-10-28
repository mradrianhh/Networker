package networker

// MessageType is deprecated.
type MessageType string

// Message types.
const (
	RESPONSE = "RESPONSE"
	REQUEST  = "REQUEST"
)

// ResponseCode ...
type ResponseCode string

// Response codes.
const (
	VALID        = "VALID"
	INVALID      = "INVALID"
	ERROR        = "ERROR"
	CONFIRMATION = "CONFIRMATION"
)

// RequestCode ...
type RequestCode string

// Request codes.
const (
	USERNAME = "USERNAME"
)
