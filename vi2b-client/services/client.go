package services

import (
	"fmt"
	"log"
	"net"
	"encoding/json"
)

type TCPServer struct {
	Conn 		net.Conn
	Address		string
	Password	string
}

type Introduce struct {
	Password	string
}

var server *TCPServer

func encodeData(dataType string, data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s`%s\n", dataType, string(jsonData))
}

func NewTCPServer(address string, password string) *TCPServer {
	server = &TCPServer{Address: address, Password: password}
	return server
}

func GetServer() *TCPServer {
	return server
}

func (s *TCPServer) Send(data string) error {
	_, err := fmt.Fprint(s.Conn, data)
	return err
}

func (s *TCPServer) Connect() {
	conn, err := net.Dial("tcp", s.Address)
	s.Conn = conn

	if err != nil {
		log.Fatal(err)
	}

	s.Send(encodeData("introduce", Introduce{Password: s.Password}))
}

func (s *TCPServer) Disconnect() {
	s.Send("bye`{}\n")
}