package main

import (
	"bufio"
	"fmt"
	"os"
)

// Task représente une tâche avec un identifiant et une description.
type Task struct {
	ID          int
	Description string
}

var tasks []Task
var nextID = 1

// Ajouter une tâche à la liste.
func addTask(description string) {
	task := Task{
		ID:          nextID,
		Description: description,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Printf("Tâche ajoutée : %d - %s\n", task.ID, task.Description)
}

// Lister toutes les tâches.
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Aucune tâche enregistrée.")
		return
	}
	fmt.Println("Liste des tâches :")
	for _, task := range tasks {
		fmt.Printf("%d - %s\n", task.ID, task.Description)
	}
}

// Supprimer une tâche par son identifiant.
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Tâche supprimée : %d\n", id)
			return
		}
	}
	fmt.Printf("Tâche non trouvée : %d\n", id)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoisissez une action :")
		fmt.Println("1. Ajouter une tâche")
		fmt.Println("2. Lister les tâches")
		fmt.Println("3. Supprimer une tâche")
		fmt.Println("4. Quitter")
		fmt.Print("> ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Entrez la description de la tâche : ")
			scanner.Scan()
			description := scanner.Text()
			addTask(description)
		case "2":
			listTasks()
		case "3":
			fmt.Print("Entrez l'ID de la tâche à supprimer : ")
			scanner.Scan()
			var id int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &id)
			if err != nil {
				fmt.Println("ID invalide.")
				continue
			}
			deleteTask(id)
		case "4":
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}