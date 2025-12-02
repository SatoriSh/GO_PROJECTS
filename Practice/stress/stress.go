package main

import (
	"fmt"
	"time"
)

var timeout bool = false
var duration int = 10
var gorutinesCount int = 25000

func main() {
	stressTimer := time.NewTimer(time.Duration(duration) * time.Second)

	fmt.Println("\nИнициализация,", gorutinesCount, "горутин...")
	for i := 0; i < gorutinesCount; i++ {
		go worker()
	}
	fmt.Println("\nГотово.")

	go func() {
		<-stressTimer.C
		timeout = true
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		seconds := 0

		for range ticker.C {
			seconds++
			fmt.Printf("\n\r%d ", seconds)

			if timeout {
				break
			}
		}
	}()

	time.Sleep(time.Duration(duration) * time.Second)
}

func worker() {
	for {
		if timeout {
			break
		}
	}
}
