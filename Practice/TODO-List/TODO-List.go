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
	showMenu()
}

func PrintAllTasks() {
	clearConsole()

	for index, task := range tasks {
		if index != 0 {
			fmt.Printf("\n%d: %s\n Terminé: %t\n", task.Id, task.Text, task.Completed)
		} else {
			fmt.Printf("%d: %s\n Terminé: %t\n", task.Id, task.Text, task.Completed)
		}
	}

	fmt.Println("\nAppuyez sur 'Entrée' pour revenir au menu")
	reader.ReadString('\n')
	showMenu()
}

func AddTask() {
	var text string
	for {
		clearConsole()

		fmt.Println("Entrez le texte de la tâche :")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur lors de l'enregistrement du texte")
			reader.ReadString('\n')
			break
		}

		text = strings.TrimSpace(input)
		if len(text) < 4 {
			fmt.Println("Erreur, veuillez entrer au moins 4 caractères")
			reader.ReadString('\n')
			continue
		}

		tasks = append(tasks, Task{Id: getMaxIdValue(), Text: text, Completed: false})

		fmt.Println("\nTâche créée avec succès, son ID :", tasks[len(tasks)-1].Id)

		fmt.Println("\nAppuyez sur 'Entrée' pour revenir au menu")
		reader.ReadString('\n')
		showMenu()
	}
}

func setCompleteStatus(status bool) {
	var id int

	for {
		clearConsole()

		fmt.Println("Entrez l'ID de la tâche que vous souhaitez marquer comme terminée")
		id = getUserValue()
		if id == -1 {
			continue
		}

		task, err := getTask(id)
		if err != nil {
			fmt.Println("\nErreur, aucune tâche avec cet ID n'existe.")
			reader.ReadString('\n')
			continue
		}

		if status {
			task.MarkComplete()
		} else {
			task.MarkNotComplete()
		}

		fmt.Println("La tâche avec l'ID", id, "a été marquée avec succès comme", task.Completed)

		fmt.Println("\nAppuyez sur 'Entrée' pour revenir au menu")
		reader.ReadString('\n')
		showMenu()
		break
	}
}

func ViewTask() {
	var id int

	for {
		clearConsole()

		fmt.Println("Entrez l'ID de la tâche que vous souhaitez consulter")
		id = getUserValue()
		if id == -1 {
			continue
		}

		task, err := getTask(id)
		if err != nil {
			fmt.Println("\nErreur, aucune tâche avec cet ID n'existe.")
			reader.ReadString('\n')
			continue
		}

		fmt.Printf("%d: %s\n Terminé: %t\n", task.Id, task.Text, task.Completed)

		fmt.Println("\nAppuyez sur 'Entrée' pour revenir au menu")
		reader.ReadString('\n')
		showMenu()
		break
	}
}

func DeleteTask() {
	var id int
	var indexInSlice int = -1

	for {
		clearConsole()

		fmt.Println("Entrez l'ID de la tâche que vous souhaitez supprimer")

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
			fmt.Println("\nErreur, aucune tâche avec cet ID n'existe. ")
			reader.ReadString('\n')
			continue
		} else {
			tasks = append(tasks[:indexInSlice], tasks[indexInSlice+1:]...)

			fmt.Printf("La tâche avec l'ID %d a été supprimée avec succès.", id)

			fmt.Println("\nAppuyez sur 'Entrée' pour revenir au menu")
			reader.ReadString('\n')
			showMenu()
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

	return &Task{}, errors.New("tâche non trouvée")
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

func showMenu() {
	for {
		clearConsole()

		fmt.Printf("\tMenu\n\n")
		fmt.Println("(1) Ajouter une tâche")
		fmt.Println("(2) Supprimer une tâche")
		fmt.Println("(3) Marquer comme terminée")
		fmt.Println("(4) Marquer comme non terminée")
		fmt.Println("(5) Voir une tâche")
		fmt.Println("(6) Voir toutes les tâches")
		fmt.Println("(7) Quitter")

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
			clearConsole()
			os.Exit(0)
		case -1: // error
			continue
		default:
			fmt.Println("Cette option n'existe pas.")
			reader.ReadString('\n')
			continue
		}
	}
}

func getUserValue() int {
	fmt.Print("\n --> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture de l'entrée :", err)
		reader.ReadString('\n')
		return -1
	}

	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erreur, veuillez entrer une valeur correcte")
		reader.ReadString('\n')
		return -1
	}

	return choice
}

func checkTasksCount() bool {
	if len(tasks) > 0 {
		return true
	} else {
		fmt.Println("\nVous n'avez aucune tâche pour effectuer cette action.")
		reader.ReadString('\n')
		return false
	}
}

func clearConsole() { fmt.Print("\033[2J\033[H") }
