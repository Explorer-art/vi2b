package core

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type Server struct {
	address		string
	password	string
}

var server *Server

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Received: %s", message)
	}
}

func NewServer(address string, password string) *Server {
	server = &Server{address: address, password: password}
	return server
}

func GetServer() *Server {
	return server
}

func (s *Server) Start() {
	http.HandleFunc("/ws", handleConnections)
	log.Printf("http server started on %s\n", s.address)
	err := http.ListenAndServe(s.address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
