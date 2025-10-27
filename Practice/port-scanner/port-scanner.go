package main

import (
	"fmt"
	"net"
	//"sync" // для синхронизированного сканирования, потому что Go слишком быстрый
)

// var wg sync.WaitGroup
var openPorts []int

func main() {
	ports := make(chan int, 100)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
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

	/*
		wg.Wait() // ждать конца выполнения горутин

		if len(openPorts) > 0 {
			sortSlace(&openPorts)
			for _, v := range openPorts {
				fmt.Println("OPEN\t", v)
			}
		}
	*/

	close(ports)
	close(results)

	sortSlice(&openPorts)
	for _, port := range openPorts {
		fmt.Println(port, "\topen")
	}
}

func worker(ports, results chan int) {
	for p := range ports { // выполняется пока канал открыт, RANGE - забирает порт из канала
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()

		results <- p
	}
}

/*
	func scanPortWithWG(index int) { // тут то же самое то с wg.Wait wg.Done wg.Add...
		defer wg.Done() // горутина завершила работу; DEFER - команда выполнится при ВЫХОДЕ из функции

		var address string = fmt.Sprintf("scanme.nmap.org:%d", index)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			return
		}
		conn.Close()

		openPorts = append(openPorts, index)
	}
*/

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
