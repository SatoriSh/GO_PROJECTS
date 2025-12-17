package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

var clients = make(map[net.Conn]bool)
var broadcast = make(chan string)
var mutex = sync.Mutex{}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on :20080")

	go handleMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting:", err)
			continue
		}

		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	clientIP := conn.RemoteAddr().String()
	broadcast <- fmt.Sprintf("Nouvel utilisateur rejoint: %s", clientIP)

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg = strings.TrimSpace(msg)

		if len(msg) > 0 {
			broadcast <- fmt.Sprintf("%s: %s", clientIP, msg)
		}
	}

	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
	broadcast <- fmt.Sprintf("L'utilisateur est sorti: %s", clientIP)
}

func handleMessages() {
	for {
		msg := <-broadcast

		mutex.Lock()
		for client := range clients {
			writer := bufio.NewWriter(client)
			writer.WriteString(msg + "\n")
			writer.Flush()
		}
		mutex.Unlock()
	}
}
