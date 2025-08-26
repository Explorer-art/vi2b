package main

import (
    "fmt"
    "net"
    "time"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
    clientAddr := conn.RemoteAddr().String()
    fmt.Printf("Новое подключение от: %s\n", clientAddr)
    
    time.Sleep(5 * time.Second)
    
    fmt.Printf("Закрытие соединения с: %s\n", clientAddr)
}