package main

import (
	"fmt"
	"os"

	"github.com/CagataySert/todo-cli/cmd"
	"github.com/CagataySert/todo-cli/internal/task"
	"github.com/CagataySert/todo-cli/storage"
)

func main() {
	taskManager := task.NewTaskManager()
	store := storage.NewStorage()
	rootCmd := cmd.NewRootCmd(store, taskManager)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
