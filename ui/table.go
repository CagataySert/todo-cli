package ui

import (
	"os"

	"github.com/aquasecurity/table"

	"github.com/CagataySert/todo-cli/internal/task"
)

func RenderTasksTable(tasks []task.Task) {
	t := table.New(os.Stdout)
	t.SetPadding(5)
	t.SetHeaderStyle(table.StyleBold)
	t.SetLineStyle(table.StyleBlue)

	t.SetHeaders("ID", "Description", "Done")

	for _, task := range tasks {
		t.AddRow(task.ID, task.Description, boolToString(task.Done))
	}

	t.Render()
}

func boolToString(b bool) string {
	if b {
		return "✔"
	}
	return "✘"
}
