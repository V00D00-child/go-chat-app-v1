package models

import (
	"encoding/json"
	"net/http"
)

// ChatMessage ...
type ChatMessage struct {
}

// MessageHandlerData ...
type MessageHandlerData struct {
	Data    interface{}
	Message string
	Type    int // 0-7
	Code    int
	Success bool
}

// SuccessHandlerMessage ...
func SuccessHandlerMessage(message string, data interface{}, w http.ResponseWriter, statusCode int, messageType int) {
	messageData := MessageHandlerData{
		Message: message,
		Data:    data,
		Type:    messageType,
		Code:    statusCode,
		Success: true,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(messageData.toJSON())
}

// FailureHandlerMessage ...
func FailureHandlerMessage(message string, w http.ResponseWriter, statusCode int, messageType int) {
	messageData := MessageHandlerData{
		Message: message,
		Data:    nil,
		Type:    messageType,
		Code:    statusCode,
		Success: false,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(messageData.toJSON())
}

func (this *MessageHandlerData) toJSON() []byte {
	data, err := json.Marshal(this)
	if err != nil {
		panic(err)
	} else {
		return data
	}
}
