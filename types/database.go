package types

import (
	"reflect"

	"go.smsk.dev/todoman/utils"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Migrations []Migration
)

type Migration struct {
	Name  string
	Model interface{}
	Data  interface{}
	Seed  interface{}
}

// to be deprecated in favor of Migration
type Seeder struct {
	Name string
	Data interface{}
}

// Migrate runs all necessary migrations
func Migrate() {
	for _, m := range Migrations {
		err := DB.AutoMigrate(m.Model)
		utils.ErrorInfo(err, "Failed to migrate %T: %v", m.Name, err)
	}
}

// Seed inserts initial data into the database
func Seed() {
	for _, m := range Migrations {
		var c int64
		DB.Model(m.Model).Count(&c)
		if c > 0 {
			// fmt.Printf("Seed data for %s already exists. Skipping...\n", m.Name)
			continue
		}
		seedsValue := reflect.ValueOf(m.Seed)
		if seedsValue.Kind() != reflect.Slice {
			if seedsValue.CanAddr() {
				seedsValue = seedsValue.Addr()
			}
			err := DB.Create(seedsValue.Interface()).Error
			utils.ErrorInfo(err, "Failed to seed %s: %v\n", m.Name, err)
			// fmt.Printf("Successfully seeded %s.\n", m.Name)
			continue
		}
		for i := 0; i < seedsValue.Len(); i++ {
			seed := seedsValue.Index(i)
			if seed.CanAddr() {
				seed = seed.Addr()
			}
			err := DB.Create(seed.Interface()).Error
			utils.ErrorInfo(err, "Failed to seed %s: %v\n", m.Name, err)
			// fmt.Printf("Successfully seeded %s.\n", m.Name)
		}
	}
}
