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
		Name:  "mark",
		Alias: []string{"m"},
		Flags: []types.Flag{
			{Name: "id", Alias: []string{"i"}, Value: ""},
			{Name: "status", Alias: []string{"s"}, Value: ""},
		},
		Description: "Mark a task as completed, If no status is provided, it will toggle the status",
		Handler:     markTask,
	})
}

// todo: needs refactoring
func markTask(flags []types.Flag) {
	_id := core.FindFlag(flags, "id")
	if _id == "" {
		_id = utils.Prompt("Task ID: ")
	}
	id, err := strconv.Atoi(_id)
	utils.Error(err)
	t := types.Task{}
	task, err := t.GetById(id)
	utils.Error(err)
	if task == nil {
		fmt.Println("Task not found")
	}
	status := core.FindFlag(flags, "status")
	if status == "" {
		if task.Completed {
			task.Completed = false
		} else {
			task.Completed = true
		}
	} else {
		if status == "done" {
			task.Completed = true
		} else {
			task.Completed = false
		}
	}
	err = task.Update()
	utils.Error(err)
	fmt.Println("Task updated successfully")
}
