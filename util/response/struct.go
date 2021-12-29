package response

type response struct {
	MessageType string      `json:"message_type,omitempty" bson:"message_type,omitempty"`
	Message     string      `json:"message,omitempty" bson:"message,omitempty"`
	Err         bool        `json:"error,omitempty" bson:"error,omitempty"`
	Data        interface{} `json:"data,omitempty" bson:"data,omitempty"`
}
