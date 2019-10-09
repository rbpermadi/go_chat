package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	connectedUsers     map[int]*User
	Messages           []*Message `json:"messages"`
	addUser            chan *User
	removeUser         chan *User
	newIncomingMessage chan *Message
	errorChannel       chan error
	doneCh             chan bool
}

func NewServer() *Server {
	Messages := []*Message{}
	connectedUsers := make(map[int]*User)
	addUser := make(chan *User)
	removeUser := make(chan *User)
	newIncomingMessage := make(chan *Message)
	errorChannel := make(chan error)
	doneCh := make(chan bool)

	return &Server{
		connectedUsers,
		Messages,
		addUser,
		removeUser,
		newIncomingMessage,
		errorChannel,
		doneCh,
	}
}

func (server *Server) AddUser(user *User) {
	log.Println("In AddUser")
	server.addUser <- user
}

func (server *Server) RemoveUser(user *User) {
	log.Println("Removing user")
	server.removeUser <- user
}

func (server *Server) ProcessNewIncomingMessage(message *Message) {
	// log.Println("In ProcessNewIncomingMessage ",message)
	server.newIncomingMessage <- message
}

func (server *Server) Done() {
	server.doneCh <- true
}

func (server *Server) sendPastMessages(user *User) {
	for _, msg := range server.Messages {
		//  log.Println("In sendPastMessages writing ",msg)
		user.Write(msg)
	}
}

func (server *Server) Err(err error) {
	server.errorChannel <- err
}

func (server *Server) SendAll(msg *Message) {
	log.Println("In Sending to all Connected users")
	for _, user := range server.connectedUsers {
		user.Write(msg)
	}
}

func (server *Server) Listen() {
	log.Println("Server Listening .....")

	for {
		select {
		// Adding a new user
		case user := <-server.addUser:
			log.Println("Added a new User")
			server.connectedUsers[user.id] = user
			log.Println("Now ", len(server.connectedUsers), " users are connected to chat room")
			server.sendPastMessages(user)

		case user := <-server.removeUser:
			log.Println("Removing user from chat room")
			delete(server.connectedUsers, user.id)

		case msg := <-server.newIncomingMessage:
			server.Messages = append(server.Messages, msg)
			server.SendAll(msg)
		case err := <-server.errorChannel:
			log.Println("Error : ", err)
		case <-server.doneCh:
			return
		}
	}

}

func (server *Server) HandleChat(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("Handling chat request ")
	conn, _ := upgrader.Upgrade(responseWriter, request, nil)

	user := NewUser(conn, server)

	log.Println("going to add user", user)
	server.AddUser(user)

	log.Println("user added successfully")
	user.Listen()

}
