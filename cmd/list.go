package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/CagataySert/todo-cli/internal/task"
	"github.com/CagataySert/todo-cli/storage"
	"github.com/CagataySert/todo-cli/ui"
)

func newListCmd(store storage.Storage, taskManager task.TaskManager) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Display all tasks in your todo list",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := store.LoadTasks()

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
				os.Exit(1)
			}

			updatedTasks := taskManager.List(tasks)

			ui.RenderTasksTable(updatedTasks)
		},
	}
}
