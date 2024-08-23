package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"go.smsk.dev/todoman/core"
	"go.smsk.dev/todoman/types"
	"go.smsk.dev/todoman/utils"
)

func init() {
	core.RegisterCommand(types.Command{
		Name:  "list",
		Alias: []string{"l"},
		Flags: []types.Flag{
			{Name: "query", Alias: []string{"q"}, Value: ""},
			{Name: "status", Alias: []string{"s"}, Value: ""},
		},
		Description: "List all tasks",
		Handler:     listTasks,
	})
}

// todo: needs refactoring
func listTasks(flags []types.Flag) {
	query := core.FindFlag(flags, "query")
	status := core.FindFlag(flags, "status")
	var tasks []types.Task
	t := types.Task{}
	result := []types.Task{}
	var err error = nil
	if query != "" {
		result, err = t.Query(query)
		utils.Error(err)
	} else if status != "" {
		result, err = t.QueryStatus(status)
		utils.Error(err)
	} else {
		result, err = t.GetAll()
		utils.Error(err)
	}
	tasks = result
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	fmt.Println("Tasks List:")
	data := [][]string{}
	for _, task := range tasks {
		status := "not done"
		if task.Completed {
			status = "done"
		}
		data = append(data, []string{strconv.Itoa(int(task.ID)), status, task.Title, task.Description})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Task Status", "Task Name", "Description"})
	table.AppendBulk(data)
	table.Render()
}
