package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

var passwdHash string
var myHash string
var charset = []byte("0123456789!$-_@#abcdefghijklmnopqrstuvwxyz")
var stop bool = false
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("\nEnter your hash:")
	fmt.Print("/> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	passwdHash = strings.TrimSpace(input)

	fmt.Println()

	maxPasswdLength := 20
	for len := 1; len <= maxPasswdLength; len++ {
		candidate := make([]byte, len)
		brute(candidate, 0, len)
		if stop {
			break
		}
	}
}

func brute(candidate []byte, index int, passwdLength int) {
	for _, ch := range charset {
		candidate[index] = ch
		//fmt.Print("\r ", string(candidate)) // TOO LONG!!!

		if index < passwdLength-1 {
			brute(candidate, index+1, passwdLength)
		}
		if stop {
			break
		}
		if hashEqual(candidate) {
			fmt.Println("\n\n", "hash", string(candidate), "=", passwdHash)
			fmt.Println("\n\n DONE! Here is your password:", string(candidate))
			stop = true
			break
		}
	}
}

func hashEqual(candidate []byte) bool {
	sum := md5.Sum(candidate)
	myHash = fmt.Sprintf("%x", sum)

	return myHash == passwdHash
}
