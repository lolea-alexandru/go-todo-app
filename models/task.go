package models

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
	Status      string `json:"status,omitempty"`
}

func ensureFileExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// Create dummy data
		tasks := Tasks{Tasks: []Task{}}

		empty_tasks_bytes, _ := json.MarshalIndent(tasks, "", "  ")

		// Create the file with empty tasks
		err = os.WriteFile(path, empty_tasks_bytes, 0664)

		if err != nil {
			fmt.Println("An error has occured: ", err)
			return
		}
	}
}

func getTasks(tasks *Tasks) {
	// Make sure the file exists before trying to read from it
	ensureFileExists("tasks.json")

	// Read the tasks from the file
	byteValue, _ := os.ReadFile("tasks.json")

	err := json.Unmarshal(byteValue, tasks)

	// Check if any errors occured during reading the file
	if err != nil {
		fmt.Println("An error has occured: ", err)
		return
	}

}

func CreateTask(Name string, Description string) {
	// Get all the tasks from the file
	tasks := Tasks{Tasks: []Task{}}
	getTasks(&tasks)

	// Create the task
	ID := uuid.New().String()
	task := Task{ID, Name, Description, "TODO"}

	tasks.Tasks = append(tasks.Tasks, task)

	tasks_bytes, _ := json.MarshalIndent(tasks, "", "  ")

	// Create the file with empty tasks
	err := os.WriteFile("tasks.json", tasks_bytes, 0664)

	if err != nil {
		fmt.Println("An error has occured: ", err)
	}
}

func ShowTasks() {
	// Get all the tasks from the file
	tasks := Tasks{Tasks: []Task{}}
	getTasks(&tasks)

	for i := 0; i < len(tasks.Tasks); i++ {
		fmt.Printf("Task #%d: %s %s\n", i+1, tasks.Tasks[i].Name, tasks.Tasks[i].Status)
		fmt.Printf("Description: %s\n\n", tasks.Tasks[i].Description)
	}
}

func UpdateTask(reader *bufio.Reader) {
	fmt.Println("Welcome to the task update menu!")

	/* ------------------- CHOOSE TASK -------------------*/
	fmt.Print("Please type in the number of the task you would like to update:")
	input_task_number, _ := reader.ReadString('\n')
	input_task_number = strings.TrimSpace(input_task_number)

	task_number, _ := strconv.Atoi(input_task_number)
	fmt.Println("You have chosen to update task:", task_number)

	/* ------------------- RETRIEVE TASK -------------------*/
	// Get all the tasks from the file
	tasks := Tasks{Tasks: []Task{}}
	getTasks(&tasks)

	/* ------------------- CHOOSE THE STATUS UPDATE -------------------*/
	fmt.Print("Please choose a new status:")
	new_status, _ := reader.ReadString('\n')
	new_status = strings.TrimSpace(new_status)

	// Update the task
	tasks.Tasks[task_number-1].Status = new_status

	tasks_bytes, _ := json.MarshalIndent(tasks, "", "  ")

	// Write to file
	err := os.WriteFile("tasks.json", tasks_bytes, 0664)

	if err != nil {
		fmt.Println("An error has occured: ", err)
	}
}

func DeleteTask(reader *bufio.Reader) {
	fmt.Println("Welcome to the task deletion menu!")

	/* ------------------- CHOOSE TASK -------------------*/
	fmt.Print("Please type in the number of the task you would like to delete:")
	input_task_number, _ := reader.ReadString('\n')
	input_task_number = strings.TrimSpace(input_task_number)

	task_number, _ := strconv.Atoi(input_task_number)
	fmt.Println("You have chosen to delete task:", task_number)

	/* ------------------- RETRIEVE TASK -------------------*/
	// Get all the tasks from the file
	tasks := Tasks{Tasks: []Task{}}
	getTasks(&tasks)

	tasks.Tasks = append(tasks.Tasks[:task_number-1], tasks.Tasks[task_number:]...)

	tasks_bytes, _ := json.MarshalIndent(tasks, "", "  ")

	// Write to file
	err := os.WriteFile("tasks.json", tasks_bytes, 0664)

	if err != nil {
		fmt.Println("An error has occured: ", err)
	}
}
