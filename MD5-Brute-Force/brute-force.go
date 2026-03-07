package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

var passwdHash string    // хеш который мы ищем
var candidateHash string // хеш текущего кандидата
var charset = []byte("0123456789!$-_@#abcdefghijklmnopqrstuvwxyz")
var stop bool = false // флаг остановки
var reader = bufio.NewReader(os.Stdin)

func main() {
	// запрашиваем хеш у пользователя
	fmt.Println("\nEnter your hash:")
	fmt.Print("/> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	passwdHash = strings.TrimSpace(input)

	// перебор паролей разной длины
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

		//fmt.Print(" ", string(candidate)) // ОЧЕНЬ ДОЛГО!!!

		// рекурсивный вызов для следующей позиции
		if index < passwdLength-1 {
			brute(candidate, index+1, passwdLength)
		}
		if stop {
			break
		}

		// проверка найденного пароля
		if hashEqual(candidate) {
			fmt.Println("\n\n", "hash", string(candidate), "=", passwdHash)
			fmt.Println("\n\n DONE! Here is your password:", string(candidate))
			stop = true
			break
		}
	}
}

func hashEqual(candidate []byte) bool {
	sum := md5.Sum(candidate)              // вычисляем MD5
	candidateHash = fmt.Sprintf("%x", sum) // преобразуем в строку

	return candidateHash == passwdHash // сравниваем
}
