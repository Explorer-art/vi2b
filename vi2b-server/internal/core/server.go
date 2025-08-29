package core

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"encoding/json"
	"strings"
	"github.com/gorilla/websocket"
)

type Server struct {
	Address		string
	Password	string
	DBName		string
}

type Client struct {
	Conn		*websocket.Conn
	IP			string
	Username	string
}

type User struct {
	ID			uint
	IP			string
	Username	string
	Is_enable	bool
}

var server *Server
var clientsData = map[*websocket.Conn]Client{}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func encodeData(dataType string, data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s`%s\n", dataType, string(jsonData))
}

func decodeData(data []byte) (string, map[string]interface{}) {
	splitedData := strings.Split(string(data), "`")

	if len(splitedData) < 2 {
		return "", nil
	}

	var jsonData map[string]interface{}
	
	json.Unmarshal([]byte(splitedData[1]), &jsonData)

	return splitedData[0], jsonData
}

func registerClient(conn *websocket.Conn, ip string, data map[string]interface{}) bool {
	if server.Password != "" {
		if data["password"] != server.Password {
			return false
		}
	}

	user, _ := DBGetUserByIP(ip)
	if user == nil {
		DBCreateUser(&User{IP: ip, Username: "", Is_enable: true})
	}

	user, _ = DBGetUserByIP(ip)
	clientsData[conn] = Client{Conn: conn, IP: ip, Username: user.Username}
	return true
}

func onData(conn *websocket.Conn, ip string, message []byte) {
	dataType, jsonData := decodeData(message)

	if dataType == "" && jsonData == nil {
		log.Println("Error decode data")
		return
	}

	if dataType == "introduce" {
		registerClient(conn, ip, jsonData)
	} else if _, ok := clientsData[conn]; ok && dataType == "bye" {
		log.Println("Bye!")
		delete(clientsData, conn)
	} else {
		for _, plugin := range plugins {
			plugin.OnMessage(conn, ip, jsonData)
		}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Received: %s", message)

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err)
		}

		onData(conn, ip, message)
	}
}

func Init(address string, password string, dbName string) *Server {
	server = &Server{Address: address, Password: password, DBName: dbName}
	DBInit(dbName)

	for name, plugin := range plugins {
		plugin.Init()
		log.Printf("Plugin %s initialized\n", name)
	}
	
	return server
}

func GetServer() *Server {
	return server
}

func (s *Server) Start() {
	http.HandleFunc("/ws", handleConnections)
	log.Printf("http server started on %s\n", s.Address)
	err := http.ListenAndServe(s.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
