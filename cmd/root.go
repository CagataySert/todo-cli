package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CagataySert/todo-cli/internal/task"
	"github.com/CagataySert/todo-cli/storage"
)

func NewRootCmd(store storage.Storage, taskManager task.TaskManager) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple command-line Todo application",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			fmt.Println("Root command is running!")
		},
	}

	rootCmd.AddCommand(newAddCmd(store, taskManager))
	rootCmd.AddCommand(newListCmd(store, taskManager))
	rootCmd.AddCommand(newMarkDoneCmd(store, taskManager))
	rootCmd.AddCommand(newDeleteCmd(store, taskManager))

	return rootCmd
}
