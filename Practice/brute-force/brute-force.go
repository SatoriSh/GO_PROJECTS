package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var passwd = []string{}
var syms = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "$", "-", "_", "@"}
var str = []string{}
var stop bool = false
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("\nEnter your password:")
	fmt.Print("/> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	input = strings.TrimSpace(input)

	for _, ch := range input {
		passwd = append(passwd, string(ch))
		str = append(str, "")
	}

	maxPasswdLength := 20

	for len := 4; len <= maxPasswdLength; len++ {
		brute(str, 0, len)
		if stop {
			break
		}
	}
}

func brute(str []string, index int, passwdLength int) {
	for i := 0; i < len(syms); i++ {
		str[index] = syms[i]
		fmt.Print(" ", str) // TOO LONG!!!

		if index < passwdLength-1 {
			brute(str, index+1, passwdLength)
		}
		if stop {
			break
		}
		if reflect.DeepEqual(str, passwd) {
			fmt.Println("\n\nDONE! Here is your password:", passwd)
			stop = true
			break
		}
	}
}
