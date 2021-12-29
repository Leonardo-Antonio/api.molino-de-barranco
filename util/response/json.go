package response

func Err(_message string, _data interface{}) *response {
	return &response{
		MessageType: "error",
		Message:     _message,
		Err:         true,
		Data:        _data,
	}
}

func Success(_message string, _data interface{}) *response {
	return &response{
		MessageType: "message",
		Message:     _message,
		Err:         false,
		Data:        _data,
	}
}
