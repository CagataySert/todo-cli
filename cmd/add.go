package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/CagataySert/todo-cli/internal/task"
	"github.com/CagataySert/todo-cli/storage"
)

func newAddCmd(store storage.Storage, taskManager task.TaskManager) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new task to your todo list",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "Error: task description is required")
				os.Exit(1)
			}

			description := strings.Join(args, " ")

			tasks := loadTasksOrExit(store)

			updatedTasks := taskManager.Add(tasks, description)

			saveTasksOrExit(store, updatedTasks)

			fmt.Println("Task added successfully!")
		},
	}
}
