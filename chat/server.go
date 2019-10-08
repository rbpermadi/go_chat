package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Messages []*Message `json: messages`
}

func NewServer() *Server {
	Messages := []*Message{}

	return &Server{
		Messages,
	}
}

func (server *Server) handlePostChat(responseWriter http.ResponseWriter, request *http.Request) *Message {
	var messageObject Message

	err := json.NewDecoder(request.Body).Decode(&messageObject)

	messageObject.Timestamp = time.Now()
	if err != nil {
		log.Println("Error while reading JSON from websocket ", err.Error())
	}

	server.Messages = append(server.Messages, &messageObject)

	return &messageObject
}
