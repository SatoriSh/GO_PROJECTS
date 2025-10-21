package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var result int
var operation string
var value1 int
var value2 int

func main() {
	calculate()
}

func calculate() {
	for {
		value1, value2 = getValues()

		calculateResult(operation)

		if operation != "!" {
			fmt.Println(value1, operation, value2, "=", result)
		} else {
			fmt.Print(value1, operation, " = ", result)
		}
		reader.ReadString('\n')
	}
}

func calculateResult(operation string) {
	switch operation {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "/":
		result = value1 / value2
	case "*":
		result = value1 * value2
	case "%":
		result = value1 % value2
	}
}

func getValues() (int, int) {
	for {
		clearConsole()

		fmt.Println("Введите первое значение")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value1, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Error")
			reader.ReadString('\n')
			continue
		}

		operation = choiceOperation()

		if operation == "!" {
			result = calculateLog(value1)
			return value1, 0
		}

		fmt.Println("Введите второе значение")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		value2, errs := strconv.Atoi(input1)

		if errs != nil {
			fmt.Println("Error")
			reader.ReadString('\n')
			continue
		}

		return value1, value2
	}
}

func choiceOperation() string {
	for {
		fmt.Println("Выберите операцию (+, -, /,*, %, !)")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "+" || input == "-" || input == "/" || input == "*" || input == "%" || input == "!" {
			return input
		} else {
			fmt.Println("Error")
			reader.ReadString('\n')
			continue
		}
	}
}

func calculateLog(a int) int {
	if a == 1 { // выход
		return a
	}

	return a * calculateLog(a-1)
}

func clearConsole() { fmt.Print("\033[2J\033[H") }
