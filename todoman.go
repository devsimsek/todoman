package main

import (
	"log"
	"os"
	"path/filepath"

	_ "go.smsk.dev/todoman/commands"
	"go.smsk.dev/todoman/core"
	"go.smsk.dev/todoman/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get home directory:", err)
	}

	// Construct the config directory path
	configDir := filepath.Join(homeDir, ".config", "todoman")

	// Check if the config directory exists, if not create it
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			log.Fatal("Failed to create config directory:", err)
		}
	}

	// Construct the database file path
	dbPath := filepath.Join(configDir, "todoman.db")

	// Connect to the database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	types.DB = db
	types.Migrate()
	types.Seed()

	if len(os.Args) < 2 {
		core.PrintHelp(core.Commands)
		return
	}
	userInput := os.Args[1]

	// Route the command
	core.MatchCommand(userInput, core.Commands)
}
