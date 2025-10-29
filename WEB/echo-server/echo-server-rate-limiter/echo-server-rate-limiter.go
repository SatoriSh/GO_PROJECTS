package main

import (
	"io"
	"log"
	"net"
)

func main() {
	connections := make(chan net.Conn, 2)

	for i := 0; i < cap(connections); i++ {
		go worker(connections) // ПАТТЕРН worker pool
	}

	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		connections <- conn
	}
}

func worker(connections chan net.Conn) { // ПАТТЕРН worker pool
	for conn := range connections {
		echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		log.Println("Writig data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
