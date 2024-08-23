package types

import (
	"gorm.io/gorm"
)

// Task struct to represent a task
type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Completed   bool   `gorm:"default:false"`
}

func init() {
	Migrations = append(Migrations, Migration{
		Name:  "Task",
		Model: &Task{},
		Seed:  []Task{},
	})
}

func (t *Task) TableName() string {
	return "tasks"
}

// GetAll fetches all tasks from the database
func (t *Task) GetAll() ([]Task, error) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// GetById fetches a single task by ID
func (t *Task) GetById(id int) (*Task, error) {
	var task Task
	result := DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

// Create inserts a new task into the database
func (t *Task) Create() error {
	result := DB.Create(t)
	return result.Error
}

// Update updates an existing task in the database
func (t *Task) Update() error {
	result := DB.Save(t)
	return result.Error
}

// Delete removes a task from the database
func (t *Task) Delete() error {
	result := DB.Delete(t)
	return result.Error
}

// Complete marks a task as completed
func (t *Task) Complete() error {
	t.Completed = true
	return t.Update()
}

// Uncomplete marks a task as incomplete
func (t *Task) Uncomplete() error {
	t.Completed = false
	return t.Update()
}

// Query fetches tasks based on a query
func (t *Task) Query(query string) ([]Task, error) {
	var tasks []Task
	result := DB.Where("title LIKE ?", "%"+query+"%").Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// Filter fetches tasks based on a status
func (t *Task) QueryStatus(status string) ([]Task, error) {
	var tasks []Task
	result := DB.Where("completed = ?", status).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// Count fetches the total number of tasks
func (t *Task) Count() (int64, error) {
	var count int64
	result := DB.Model(&Task{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
