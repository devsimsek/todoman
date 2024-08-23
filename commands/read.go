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
		Name:  "read",
		Alias: []string{"r"},
		Flags: []types.Flag{
			{Name: "id", Alias: []string{"i"}, Value: ""},
		},
		Description: "Read a task by ID",
		Handler:     readTask,
	})
}

// todo: needs refactoring
func readTask(flags []types.Flag) {
	_id := core.FindFlag(flags, "id")
	if _id == "" {
		_id = utils.Prompt("Task ID: ")
	}
	id, err := strconv.Atoi(_id)
	utils.Error(err)
	task := types.Task{}
	result, err := task.GetById(id)
	utils.Error(err)
	if result == nil {
		fmt.Println("Task not found")
		return
	}
	status := "not done"
	if result.Completed {
		status = "done"
	}
	data := [][]string{
		{strconv.Itoa(int(result.ID)), status, result.Title, result.Description},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Task Status", "Task Name", "Description"})
	table.AppendBulk(data)
	table.Render()
}
