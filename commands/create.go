package commands

import (
	"fmt"

	"go.smsk.dev/todoman/core"
	"go.smsk.dev/todoman/types"
	"go.smsk.dev/todoman/utils"
)

func init() {
	core.RegisterCommand(types.Command{
		Name:  "create",
		Alias: []string{"c"},
		Flags: []types.Flag{
			{Name: "title", Alias: []string{"t"}, Value: ""},
			{Name: "description", Alias: []string{"d"}, Value: ""},
		},
		Description: "Create a new task",
		Handler:     createTask,
	})
}

func createTask(flags []types.Flag) {
	title := ""
	description := ""
	if core.FindFlag(flags, "title") != "" || core.FindFlag(flags, "description") != "" {
		title = core.FindFlag(flags, "title")
		description = core.FindFlag(flags, "description")
	} else {
		title = utils.Prompt("Task title: ")
		description = utils.Prompt("Task description: ")
	}
	task := types.Task{
		Title:       title,
		Description: description,
	}
	result := types.DB.Create(&task)
	utils.Error(result.Error)
	fmt.Println("Task created successfully")
}
