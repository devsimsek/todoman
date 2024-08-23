package commands

import (
	"fmt"
	"strconv"

	"go.smsk.dev/todoman/core"
	"go.smsk.dev/todoman/types"
	"go.smsk.dev/todoman/utils"
)

func init() {
	core.RegisterCommand(types.Command{
		Name:  "delete",
		Alias: []string{"d"},
		Flags: []types.Flag{
			{Name: "id", Alias: []string{"i"}, Value: ""},
		},
		Description: "Delete a task by ID",
		Handler:     deleteTask,
	})
}

// todo: needs refactoring
func deleteTask(flags []types.Flag) {
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

	err = result.Delete()
	utils.Error(err)
	fmt.Println("Task deleted successfully")
}
