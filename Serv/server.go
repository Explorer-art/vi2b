package main

import (
    "fmt"
    "net"
)

type Server struct {
    address string
}

func NewServer(address string) *Server {
    return &Server{address: address}
}

func (s *Server) Start() {
    listener, err := net.Listen("tcp", s.address)
    if err != nil {
        fmt.Printf("Ошибка запуска сервера: %v\n", err)
        return
    }
    defer listener.Close()
    
    fmt.Printf("Сервер запущен на %s\n", s.address)
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Ошибка принятия соединения: %v\n", err)
            continue
        }
        
        go handleConnection(conn)
    }
}