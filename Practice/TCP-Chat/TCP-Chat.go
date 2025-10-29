package main

import (
	"bufio"
	"log"
	"net"
)

var uniqueAddr net.Addr

func connWorker(connections chan net.Conn, messages chan string) {
	for conn := range connections {
		for {
			reader := bufio.NewReader(conn)

			s, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Unable to read data")
			}

			uniqueAddr = conn.RemoteAddr()
			messages <- s
		}
	}
}

func msgWorker(connSlice *[]net.Conn, messages chan string) {
	for {
		for msg := range messages {
			for _, conn := range *connSlice {
				if conn.RemoteAddr() == uniqueAddr {
					continue
				}

				writer := bufio.NewWriter(conn)
				if _, err := writer.WriteString(msg); err != nil {
					log.Println("Unable to write data")
				}

				writer.Flush()
			}
		}
	}
}

func main() {
	connections := make(chan net.Conn, 3)
	messages := make(chan string)

	connSlice := make([]net.Conn, 0, cap(connections))

	for i := 0; i < cap(connections); i++ {
		go connWorker(connections, messages)
	}
	go msgWorker(&connSlice, messages)

	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Error", err.Error())
	}

	log.Println("Listening on: 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error", err.Error())
		}

		if len(connSlice) < cap(connections) {
			connSlice = append(connSlice, conn)
		}
		
		connections <- conn
	}
}
