package main

import (
	"bufio"
	"log"
	"net"
)

var senderAddr net.Addr

func connWorker(connections chan net.Conn, messages chan string, connSlice *[]net.Conn) {
	for conn := range connections {
		reader := bufio.NewReader(conn)

		for {
			s, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Client disconnected")
				removeConn(connSlice, conn)
				break
			}

			senderAddr = conn.RemoteAddr()
			s = senderAddr.String() + ": " + s
			s = s[6:]
			messages <- s
		}
	}
}

func removeConn(connSlice *[]net.Conn, conn net.Conn) {
	for i, value := range *connSlice {
		if value == conn {
			*connSlice = append((*connSlice)[:i], (*connSlice)[i+1:]...)
		}
	}
}

func msgWorker(connSlice *[]net.Conn, messages chan string) {
	for {
		for msg := range messages {
			for _, conn := range *connSlice {
				if conn.RemoteAddr() == senderAddr {
					continue
				}

				writer := bufio.NewWriter(conn)
				if _, err := writer.WriteString("\n" + msg + "\n"); err != nil {
					log.Println("Unable to write data")
				}
				writer.Flush()
			}
		}
	}
}

func queueWorker(conn net.Conn, connections chan net.Conn, connSlice *[]net.Conn) {
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString("\nThere is no place, wait\n\n"); err != nil {
		log.Println("Unable to write data")
	}
	writer.Flush()

	for {
		if len(*connSlice) < cap(connections) {
			*connSlice = append(*connSlice, conn)
			connections <- conn

			if _, err := writer.WriteString("\nThe place is free, you can write\n\n"); err != nil {
				log.Println("Unable to write data")
			}
			writer.Flush()
			break
		}
	}
}

func main() {
	connections := make(chan net.Conn, 2)
	messages := make(chan string)

	connSlice := make([]net.Conn, 0, cap(connections))

	for i := 0; i < cap(connections); i++ {
		go connWorker(connections, messages, &connSlice)
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
			connections <- conn
		} else {
			go queueWorker(conn, connections, &connSlice)
		}
	}
}
