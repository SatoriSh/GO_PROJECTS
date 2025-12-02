package main

import (
	"fmt"
	"time"
)

var timeout bool = false
var duration int = 10
var gorutinesCount int = 200000

func main() {
	fmt.Println("\nИнициализация...")
	fmt.Println()
	for i := 0; i < gorutinesCount; i++ {
		go worker()
		fmt.Printf("\rСоздано горутин: %d/%d", i+1, gorutinesCount)
	}
	fmt.Println("\nГотово.")

	start := time.Now().Round(time.Second)
	stressTimer := time.NewTimer(time.Duration(duration) * time.Second)
	go func() {
		<-stressTimer.C
		timeout = true
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		fmt.Println()
		for range ticker.C {
			fmt.Printf("\r%v ", time.Since(start))

			if timeout {
				break
			}
		}
	}()

	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("\r%v ", time.Since(start))
}

func worker() {
	for {
		if timeout {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
