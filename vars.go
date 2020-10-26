package networker

// MessageType is a string representation of the type of message. I.E "AUTHENTICATION", "CONFIRMATION", etc.
type MessageType string

// Message types.
const (
	Confirmation   = "CONFIRMATION"
	Authentication = "AUTHENTICATION"
	Event          = "EVENT"
	Registration   = "REGISTRATION"
	Identification = "IDENTIFICATION"
)

// ResponseCode ...
type ResponseCode string

// Response codes
const (
	Accepted    = "ACCEPTED"
	NotAccepted = "NOT ACCEPTED"
)

// Events
const (
	LoggedOut = "LOGGED OUT"
)
