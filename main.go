package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rbpermadi/go_chat/chat"
	"github.com/rbpermadi/go_chat/config"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

var configuration config.Configuration
var serverHostName string

func init() {
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/rbpermadi/go_chat/.env")

	configuration = config.LoadConfigAndSetUpLogging()

	serverHostName = fmt.Sprintf("%s:%s", configuration.HostName, strconv.Itoa(configuration.Port))

	log.Println("The serverHost url", serverHostName)

}

func main() {

	server := chat.NewServer()
	go server.Listen()

	router := httprouter.New()
	router.GET("/messages", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(server)
	})

	router.POST("/messages", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var messageObject chat.Message

		err := json.NewDecoder(r.Body).Decode(&messageObject)

		messageObject.CreatedAt = time.Now()
		if err != nil {
			log.Println("Error while reading JSON from websocket ", err.Error())
		}

		server.Messages = append(server.Messages, &messageObject)
		server.SendAll(&messageObject)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(messageObject)
	})

	router.GET("/chat", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		server.HandleChat(w, r)
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "index.html")
	})

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})
	http.ListenAndServe(serverHostName, co.Handler(router))

}
