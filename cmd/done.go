package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/CagataySert/todo-cli/internal/task"
	"github.com/CagataySert/todo-cli/storage"
)

func newMarkDoneCmd(store storage.Storage, taskManager task.TaskManager) *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "Mark a task as completed",
		Run: func(cmd *cobra.Command, args []string) {
			id := requireSingleArg(args)
			tasks := loadTasksOrExit(store)

			updatedTasks, err := taskManager.MarkDone(tasks, id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			saveTasksOrExit(store, updatedTasks)

			fmt.Println("Task updated successfully!")
		},
	}
}
