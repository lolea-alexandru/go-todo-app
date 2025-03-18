package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lolea-alexandru/go-todo-app/models"
)

func createTask(reader *bufio.Reader) {
	fmt.Println("Welcome to the task creation menu!")

	// Get the task name
	fmt.Print("Please enter the task name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Get the task name
	fmt.Print("Please enter the task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	// Create task
	models.CreateTask(name, description)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to your TODO app!\n")
	fmt.Print("What would you like to do?\n")
	for true {
		// Print the menu
		fmt.Print("1. Create a task\n")
		fmt.Print("2. Get all tasks\n")
		fmt.Print("3. Update a task\n")
		fmt.Print("4. Delete a task\n")

		fmt.Print("Please input the number corresponding to your option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			createTask(reader)
		case "2":
			models.ShowTasks()
		case "3":
			fmt.Println("You chose to update a task")

		case "4":
			fmt.Println("You chose delete a task")

		default:
			fmt.Println("Invalid option, please try again")
		}
	}

	// models.CreateTask(1, "Do Dishes", "Wash the dishes");
}
