package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var tasks []Task // срез задач

type Task struct {
	Id        int
	Text      string
	Completed bool
}

func (t *Task) MarkComplete()    { t.Completed = true }
func (t *Task) MarkNotComplete() { t.Completed = false }

func main() {
	getMenu()
}

func UpdateText(t *Task) {
	fmt.Println("Введите новый текст для", t.Id)
	input1, _ := reader.ReadString('\n')
	t.Text = input1
}

func PrintAllTasks() {
	clearConsole()

	for index, task := range tasks {
		if index != 0 {
			fmt.Printf("\n%d: %s\n Completed: %t\n", task.Id, task.Text, task.Completed)
		} else {
			fmt.Printf("%d: %s\n Completed: %t\n", task.Id, task.Text, task.Completed)
		}
	}

	fmt.Println("\nНажмите'Enter' чтобы вернуться в меню")
	reader.ReadString('\n')
	getMenu()
}

func AddTask() {
	var text string
	for {
		clearConsole()

		fmt.Println("Введите текст задачи:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка сохранения текста")
			reader.ReadString('\n')
			break
		}

		text = strings.TrimSpace(input)
		if len(text) < 4 {
			fmt.Println("Ошибка, введите минимум 4 символа")
			reader.ReadString('\n')
			continue
		}

		tasks = append(tasks, Task{Id: getMaxIdValue(), Text: text, Completed: false})

		fmt.Println("\nЗадача успешно создана, её ID:", tasks[len(tasks)-1].Id)

		fmt.Println("\nНажмите'Enter' чтобы вернуться в меню")
		reader.ReadString('\n')
		getMenu()
	}
}

func setCompleteStatus(status bool) {
	var id int

	for {
		clearConsole()

		fmt.Println("Введите ID задачи которую хотите пометить как выполненную")
		id = getUserValue()
		if id == -1 {
			continue
		}

		task, err := getTask(id)
		if err != nil {
			fmt.Println("\nОшибка, задачи с таким ID не существует.")
			reader.ReadString('\n')
			continue
		}

		if status {
			task.MarkComplete()
		} else {
			task.MarkNotComplete()
		}

		fmt.Println("Задача с ID", id, "была успешно помечена как", task.Completed) // почемуто здесь всегда false

		fmt.Println("\nНажмите'Enter' чтобы вернуться в меню")
		reader.ReadString('\n')
		getMenu()
		break
	}
}

func ViewTask() {
	var id int

	for {
		clearConsole()

		fmt.Println("Введите ID задачи которую хотите просмотреть")
		id = getUserValue()
		if id == -1 {
			continue
		}

		task, err := getTask(id)
		if err != nil {
			fmt.Println("\nОшибка, задачи с таким ID не существует.")
			reader.ReadString('\n')
			continue
		}

		fmt.Printf("%d: %s\n Completed: %t\n", task.Id, task.Text, task.Completed)

		fmt.Println("\nНажмите'Enter' чтобы вернуться в меню")
		reader.ReadString('\n')
		getMenu()
		break
	}
}

func DeleteTask() {
	var id int
	var indexInSlice int = -1

	for {
		clearConsole()

		fmt.Println("Введите ID задачи которую хотите удалить")

		id = getUserValue()

		if id == -1 {
			continue
		}

		for index, task := range tasks {
			if task.Id == id {
				indexInSlice = index
			}
		}

		if indexInSlice == -1 {
			fmt.Println("\nОшибка, задачи с таким ID не существует. ")
			reader.ReadString('\n')
			continue
		} else {
			tasks = append(tasks[:indexInSlice], tasks[indexInSlice+1:]...)

			fmt.Printf("Задача с id %d была успешно удалена.", id)
			orderTasks()

			fmt.Println("\nНажмите'Enter' чтобы вернуться в меню")
			reader.ReadString('\n')
			getMenu()
			break
		}
	}
}

func getTask(id int) (*Task, error) { // *Task - возвращаем указатель на ориг
	for index := range tasks {
		if tasks[index].Id == id {
			return &tasks[index], nil
		}
	}

	return &Task{}, errors.New("Task not found")
}

func orderTasks() {
	for i := 0; i < len(tasks); i++ {
		tasks[i].Id = i + 1
	}
}

func getMaxIdValue() int {
	var maxIdValue int

	for _, task := range tasks {
		if task.Id > maxIdValue {
			maxIdValue = task.Id
		}
	}
	return maxIdValue + 1
}

func getMenu() {
	for {
		clearConsole()

		fmt.Printf("Меню\n\n")
		fmt.Println("(1) Добавить задачу")
		fmt.Println("(2) Удалить задачу")
		fmt.Println("(3) Пометить как выполненную")
		fmt.Println("(4) Пометить как не выполненную")
		fmt.Println("(5) Посмотреть задачу")
		fmt.Println("(6) Посмотреть все задачи")
		fmt.Println("(7) Выйти")

		switch getUserValue() {
		case 1:
			AddTask()
		case 2:
			if checkTasksCount() {
				DeleteTask()
			}
		case 3:
			if checkTasksCount() {
				setCompleteStatus(true)
			}
		case 4:
			if checkTasksCount() {
				setCompleteStatus(false)
			}
		case 5:
			if checkTasksCount() {
				ViewTask()
			}
		case 6:
			if checkTasksCount() {
				PrintAllTasks()
			}
		case 7:
			os.Exit(0)
		case -1: // error
			continue
		default:
			fmt.Println("Такой опции не существует.")
			reader.ReadString('\n')
			continue
		}
	}
}

func getUserValue() int {

	fmt.Print("\n --> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		reader.ReadString('\n')
		return -1
	}

	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести число")
		reader.ReadString('\n')
		return -1
	}

	return choice
}

func checkTasksCount() bool {
	if len(tasks) > 0 {
		return true
	} else {
		fmt.Println("\nУ вас нет ни одной задачи чтобы выполнить это действие.")
		reader.ReadString('\n')
		return false
	}
}

func clearConsole() { fmt.Print("\033[2J\033[H") }
