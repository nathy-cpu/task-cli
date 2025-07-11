package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	StorageFile = "tasks.json"
	TaskTodo    = "todo"
	TaskPending = "pending"
	TaskDone    = "done"
)

type Task struct {
	Id          uint      `json: "id"`
	Description string    `json: "description"`
	Status      string    `json: "status"`
	CreatedAt   time.Time `json: "createdAt"`
	UpdatedAt   time.Time `json: "updatedAt"`
}

func AddTask(id uint, description string, status string) error {
	task := Task{
		Id:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	data, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(StorageFile, data, 0644)
}

func CheckFile() {
	fmt.Println("Checking storage file ...")
	_, err := os.Stat(StorageFile)
	if os.IsNotExist(err) {
		fmt.Println("Storage file does not exist. Creating storage file ...")
		file, err := os.Create(StorageFile)
		if err != nil {
			fmt.Println("Failed to create storage file: ", err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Println("Storage file created successfully")
	} else {
		fmt.Println("Storage file found")
	}
}
