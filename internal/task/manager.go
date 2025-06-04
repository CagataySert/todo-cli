package task

import (
	"errors"
	"slices"

	"github.com/gofrs/uuid"
)

type TaskManager struct{}

func NewTaskManager() TaskManager {
	return TaskManager{}
}

func (tm *TaskManager) Add(tasks []Task, description string) []Task {
	id, _ := uuid.NewV4()

	newTask := Task{
		ID:          id.String(),
		Description: description,
		Done:        false,
	}

	return append(tasks, newTask)
}

func (tm *TaskManager) List(tasks []Task) []Task {
	return tasks
}

func (tm *TaskManager) MarkDone(tasks []Task, id string) ([]Task, error) {
	index, err := findTaskIndex(tasks, id)
	if err != nil {
		return nil, err
	}
	tasks[index].Done = true

	return tasks, nil
}

func (tm *TaskManager) Delete(tasks []Task, id string) ([]Task, error) {
	index, err := findTaskIndex(tasks, id)
	if err != nil {
		return nil, err
	}
	return slices.Delete(tasks, index, index+1), nil
}

func findTaskIndex(tasks []Task, id string) (int, error) {
	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.ID == id
	})

	if index == -1 {
		return -1, errors.New("task not found with id: " + id)
	}
	return index, nil
}
