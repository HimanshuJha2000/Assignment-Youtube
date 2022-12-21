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
	client = connect(config.GetConfig().Database)
}

func Client() *gorm.DB {
	return client
}

func connect(database config.Database) *gorm.DB {
	var err error

	db, err := gorm.Open(database.Dialect, database.URL())

	if err != nil {
		panic("database connection failure")
	}

	// This will prevent update or delete without where clause
	db.BlockGlobalUpdate(true)

	// setting db connection params
	db.DB().SetMaxIdleConns(database.MaxIdleConnections)
	db.DB().SetMaxOpenConns(database.MaxOpenConnections)

	return db
}
