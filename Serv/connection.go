package main

import (
    "encoding/json"
    "fmt"
    "net"
    "time"
)

type ServerMessage struct {
    DataType string `json:"dataType"`
    Message  string `json:"message"`
    ClientIP string `json:"clientIP"`
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    clientAddr := conn.RemoteAddr().String()
    fmt.Printf("Новое подключение от: %s\n", clientAddr)
    
    clientIP := conn.RemoteAddr().(*net.TCPAddr).IP.String()
    
    welcomeMsg := ServerMessage{
        DataType: "connection",
        Message:  "Добро пожаловать на сервер!",
        ClientIP: clientIP,
    }
    
    jsonData, err := json.Marshal(welcomeMsg)
    if err != nil {
        fmt.Printf("Ошибка создания JSON: %v\n", err)
        return
    }
    
    _, err = conn.Write(jsonData)
    if err != nil {
        fmt.Printf("Ошибка отправки данных: %v\n", err)
        return
    }
    
    fmt.Printf("Отправлено приветственное сообщение клиенту %s\n", clientAddr)
    
    // Ждем 5 секунд
    time.Sleep(5 * time.Second)

    closeMsg := ServerMessage{
        DataType: "disconnection", 
        Message:  "Соединение закрыто",
        ClientIP: clientIP,
    }
    
    jsonData, err = json.Marshal(closeMsg)
    if err != nil {
        fmt.Printf("Ошибка создания JSON: %v\n", err)
        return
    }
    
    conn.Write(jsonData)
    
    fmt.Printf("Закрытие соединения с: %s\n", clientAddr)
}