package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CagataySert/todo-cli/internal/task"
)

type Storage interface {
	SaveTasks(tasks []task.Task) error
	LoadTasks() ([]task.Task, error)
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) SaveTasks(tasks []task.Task) error {
	file, err := os.Create("tasks.json")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	return nil
}

func (s *storage) LoadTasks() ([]task.Task, error) {
	file, err := os.Open("tasks.json")

	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}

		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	var tasks []task.Task

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("failed to decode tasks: %w", err)
	}

	return tasks, nil
}
