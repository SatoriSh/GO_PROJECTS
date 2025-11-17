package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"unicode/utf8"
)

var openPorts []int
var reader = bufio.NewReader(os.Stdin)

func main() {
	ports := make(chan int, 300)
	results := make(chan int)
	address := getAddress()

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sortSlice(&openPorts)
	for _, port := range openPorts {
		fmt.Println(port, "\topen")
	}
}

func worker(ports, results chan int, addressToScan string) {
	for p := range ports { // выполняется пока канал открыт, RANGE - забирает порт из канала
		address := net.JoinHostPort(addressToScan, fmt.Sprintf("%d", p))
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()

		results <- p
	}
}

func getAddress() string {
	for {
		fmt.Println("Enter the address to scan")

		fmt.Print("--> ")
		input, err := reader.ReadString('\n')
		if str := utf8.RuneCountInString(input); err != nil || str < 6 {
			fmt.Println(err)
			continue
		}

		input = strings.TrimSpace(input)
		return input
	}
}

func sortSlice(slice *[]int) {
	var value2 int

	for i := 0; i < len(*slice)-1; i++ {
		for j := 0; j < len(*slice)-1; j++ {
			if (*slice)[j] > (*slice)[j+1] {
				value2 = (*slice)[j+1]
				(*slice)[j+1] = (*slice)[j]
				(*slice)[j] = value2
			}
		}
	}
}
