package core

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn 		*websocket.Conn
	Address		string
	Password	string
	Username	string
}

type Introduce struct {
	Password	string
}

var client *Client

func encodeData(dataType string, data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s`%s\n", dataType, string(jsonData))
}

func NewClient(address string, password string) *Client {
	client = &Client{Address: address, Password: password}
	return client
}

func GetClient() *Client {
	return client
}

func (c *Client) Send(message []byte) error {
	err := c.Conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Printf("Write error: %s\n", err)
	}
	return err
}

func (c *Client) Connect() error {
	url := "ws://" + c.Address + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	c.Conn = conn

	c.Send([]byte(encodeData("introduce", Introduce{Password: c.Password})))

	return nil
}

func (c *Client) Disconnect() {
	c.Send([]byte("bye`{}\n"))
	c.Conn.Close()
}