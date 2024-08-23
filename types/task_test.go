package types_test

import (
	"testing"

	"go.smsk.dev/todoman/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Testing task.go
func init() {
	// set DB connection
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	types.DB = db

	types.Migrate()
}

// Task instance
var task types.Task

// TestCreate function
func TestCreate(t *testing.T) {
	task := types.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}

	if err := task.Create(); err != nil {
		t.Errorf("Error creating task: %v", err)
	}

	t.Logf("Created task: %v", task)
}

// TestGetAll function
func TestGetAll(t *testing.T) {
	tasks, err := task.GetAll()
	if err != nil {
		t.Errorf("Error fetching tasks: %v", err)
	}

	if len(tasks) == 0 {
		t.Errorf("Expected tasks, got none")
	}

	if len(tasks) > 0 {
		t.Logf("Fetched %d tasks", len(tasks))
	}
}

// TestGetById function
func TestGetById(t *testing.T) {
	task, err := task.GetById(1)
	if err != nil {
		t.Errorf("Error fetching task: %v", err)
	}

	if task == nil {
		t.Errorf("Expected task, got none")
	}

	if task != nil {
		t.Logf("Fetched task: %v", task)
	}
}

// TestUpdate function
func TestUpdate(t *testing.T) {
	task, err := task.GetById(1)
	if err != nil {
		t.Errorf("Error fetching task: %v", err)
	}

	task.Completed = true
	if err := task.Update(); err != nil {
		t.Errorf("Error updating task: %v", err)
	}

	t.Logf("Updated task: %v", task)
}

// TestComplete function
func TestComplete(t *testing.T) {
	task, err := task.GetById(1)
	if err != nil {
		t.Errorf("Error fetching task: %v", err)
	}

	if err := task.Complete(); err != nil {
		t.Errorf("Error completing task: %v", err)
	}

	t.Logf("Completed task: %v", task)
}

// TestDelete function
func TestDelete(t *testing.T) {
	task, err := task.GetById(1)
	if err != nil {
		t.Errorf("Error fetching task: %v", err)
	}
	if err := task.Delete(); err != nil {
		t.Errorf("Error deleting task: %v", err)
	}

	t.Logf("Deleted task: %v", task)
}
