package config

import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Category struct {
	Id   string
	Name string
}

func (Category) TableName() string {
	return "category"
}

func ConnectDB() (*gorm.DB, error) {
	// return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	connectionString := fmt.Sprintf("root:Admin123@tcp(db:3306)/mydb")

	var err error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}

	return db, err
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
