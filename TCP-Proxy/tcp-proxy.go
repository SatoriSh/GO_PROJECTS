package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatalln("Failed to connect to destination")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln("Forward error:", err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln("Response error:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Listen failed")
	}

	log.Println("Listening on 0.0.0.0:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Accept failed:", err)
		}

		go handle(conn)
	}
}
