package helper

const (
	MESSAGE = "message"
	ERROR   = "error"
)

type Response struct {
	MessageType string      `json:"message_type" bson:"message_type" xml:"message_type"`
	Message     string      `json:"message" bson:"message" xml:"message"`
	Error       bool        `json:"error" bson:"error" xml:"error"`
	Data        interface{} `json:"data" bson:"data" xml:"data"`
}
