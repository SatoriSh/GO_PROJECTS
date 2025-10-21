package main

import (
	"fmt"
	"net"
	"sync" // для синхронизированного сканирования, потому что Go слишком быстрый
)

var wg sync.WaitGroup
var openPorts []int

func main() {
	for i := 1; i < 1024; i++ {
		wg.Add(1) // зарегистрировать новую горутину
		go scanPort(i)
	}

	wg.Wait() // ждать конца выполнения горутины (вызывается один раз после ВСЕХ горутин)

	sortSlace(&openPorts)
	for _, v := range openPorts {
		fmt.Println("OPEN\t", v)
	}
}

func scanPort(index int) {
	defer wg.Done() // горутина завершила работу; DEFER - команда выполнится при ВЫХОДЕ из функции

	var address string = fmt.Sprintf("scanme.nmap.org:%d", index)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	conn.Close()

	openPorts = append(openPorts, index)
}

func sortSlace(slice *[]int) {
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
