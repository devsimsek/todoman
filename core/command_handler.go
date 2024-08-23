package core

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go.smsk.dev/todoman/types"
	"go.smsk.dev/todoman/utils"
)

var Commands = []types.Command{}

func RegisterCommand(cmd types.Command) {
	Commands = append(Commands, cmd)
}

// PrintHelp function to print the help menu
func PrintHelp(commands []types.Command) {
	fmt.Println("Usage: todoman [command]")
	fmt.Println("Commands:")
	for _, cmd := range commands {
		fmt.Printf("\t%s\n\t\t%s\n", cmd.Name, cmd.Description)
		if len(cmd.Flags) > 0 {
			fmt.Printf("\tFlags:\n")
			for _, f := range cmd.Flags {
				fmt.Printf("\t  -%s\t%s\n", f.Name, f.Value)
			}
		}
		if len(cmd.Alias) > 0 {
			fmt.Printf("\tAliases: %s\n", strings.Join(cmd.Alias, ", "))
		}
	}
}

// FindFlag function to find the value of a flag in a slice of flags
func FindFlag(flags []types.Flag, name string) string {
	for _, f := range flags {
		if f.Name == name {
			if f.Value == "nil" {
				return ""
			}
			return f.Value
		}
	}

	return ""
}

func MatchCommand(input string, commands []types.Command) {
	// Add the help command
	commands = append(commands, types.Command{
		Name:        "help",
		Alias:       []string{"h"},
		Description: "Show available commands",
		Handler: func(flags []types.Flag) {
			PrintHelp(commands)
		},
	})

	for _, cmd := range commands {
		if input == cmd.Name || utils.Contains(cmd.Alias, input) {
			cmdFlags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
			for i, f := range cmd.Flags {
				cmdFlags.StringVar(&cmd.Flags[i].Value, f.Name, f.Value, "")
				if f.Alias != nil {
					for _, a := range f.Alias {
						cmdFlags.StringVar(&cmd.Flags[i].Value, a, f.Value, "")
					}
				}
			}

			if err := cmdFlags.Parse(os.Args[2:]); err != nil {
				log.Fatal(err)
			}

			cmd.Handler(cmd.Flags)
			return
		}
	}

	// If no matching command is found, print an error message
	fmt.Println("Command not found. Use 'todoman help' to see available commands.")
}
