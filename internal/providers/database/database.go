package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/razorpay/MachineRound/internal/config"
)

var (
	client *gorm.DB
)

func Initialize() {
	fmt.Println("Initializing postgres database...")
	client = Connect(config.GetConfig().Database)
}

func Client() *gorm.DB {
	return client
}

func Connect(database config.Database) *gorm.DB {
	var err error

	db, err := gorm.Open(database.Dialect, database.URL())

	if err != nil {
		panic("database connection failure")
	}

	// This will prevent update or delete without where clause
	db.BlockGlobalUpdate(true)

	return db
}
