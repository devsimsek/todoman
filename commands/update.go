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
		Name:  "update",
		Alias: []string{"u"},
		Flags: []types.Flag{
			{Name: "id", Alias: []string{"i"}, Value: ""},
			{Name: "title", Alias: []string{"t"}, Value: ""},
			{Name: "description", Alias: []string{"d"}, Value: ""},
			{Name: "status", Alias: []string{"s"}, Value: ""},
		},
		Description: "Update a task by ID",
		Handler:     updateTask,
	})
}

// todo: needs refactoring
func updateTask(flags []types.Flag) {
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

	title := core.FindFlag(flags, "title")
	description := core.FindFlag(flags, "description")
	status := core.FindFlag(flags, "status")

	if title == "" {
		title = utils.Prompt("Task title (leave empty to skip): ")
	}
	if description == "" {
		description = utils.Prompt("Task description (leave empty to skip): ")
	}
	if status == "" {
		status = utils.Prompt("Task status (leave empty to skip): ")
	}

	if title != "" {
		result.Title = title
	}
	if description != "" {
		result.Description = description
	}
	if status != "" {
		if status == "done" {
			result.Completed = true
		} else {
			result.Completed = false
		}
	}
	err = result.Update()
	utils.Error(err)
	fmt.Println("Task updated successfully")
}
