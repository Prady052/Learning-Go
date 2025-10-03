package db

import (
	"fmt"
	//"log"
	//"time"

	"cdac.com/day5/models"
	//go get go.uber.org/zap
	//go get gorm.io/driver/postgres
	//"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect establishes GORM connection and configures pool
func Connect() {

	// creating the connection string
	//dsn (Data Source Name)
	dsn := "host=localhost user=postgres password=root123 dbname=Employees port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Printf("failed to connect to database: %v", err)
	}

	//if err != nil {
	//	fmt.Printf("failed to get sql.DB from gorm: %v", err)
	//}

	// GORM automigrate as a convenience (still use migrations in production)
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("auto-migrate warning: %v", err)
	}

	DB = db
	fmt.Println("connected to database")
}
