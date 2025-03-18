package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

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
