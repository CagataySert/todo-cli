package cmd

import (
	"fmt"
	"os"

	"github.com/CagataySert/todo-cli/internal/task"
	"github.com/CagataySert/todo-cli/storage"
)

func requireSingleArg(args []string) string {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: task id is required")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Error: only one task can be provided")
		os.Exit(1)
	}
	return args[0]
}

func loadTasksOrExit(store storage.Storage) []task.Task {
	tasks, err := store.LoadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}
	return tasks
}

func saveTasksOrExit(store storage.Storage, tasks []task.Task) {
	if err := store.SaveTasks(tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		os.Exit(1)
	}
}
