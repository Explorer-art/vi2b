package core

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/chzyer/readline"
	"github.com/gorilla/websocket"
)

type Server struct {
	Conn 		*websocket.Conn
	Address		string
	Password	string
}

type Introduce struct {
	Password	string `json:"password"`
}

type SendCommand struct {
	Command 	string `json:"command"`
}

var server *Server
var rl *readline.Instance

func encodeData(dataType string, data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s`%s\n", dataType, string(jsonData))
}

func Init(l *readline.Instance) {
	rl = l
}

func NewServer(address string, password string) *Server {
	server = &Server{Address: address, Password: password}
	return server
}

func GetServer() *Server {
	return server
}

func (s *Server) Send(message []byte) error {
	err := s.Conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Printf("Write error: %s\n", err)
	}
	return err
}

func (s *Server) SendMessage(dataType string, data interface{}) error {
	err := s.Conn.WriteMessage(websocket.TextMessage, []byte(encodeData(dataType, data)))
	if err != nil {
		log.Printf("Write error: %s\n", err)
	}
	return err
}

func (s *Server) OnData() {
	defer s.Conn.Close()

	for {
		_, message, err := s.Conn.ReadMessage()
		if err != nil {
			log.Fatal("Received error: ", err)
			return
		}

		rl.Write([]byte("\rCHAT > " + string(message) + "\n"))
		rl.Refresh()
	}
}

func (s *Server) Connect() error {
	url := "ws://" + s.Address + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.Conn = conn

	go s.OnData()

	s.SendMessage("introduce", Introduce{Password: s.Password})

	return nil
}

func (s *Server) Disconnect() {
	s.Send([]byte("bye`{}\n"))
	s.Conn.Close()
}
