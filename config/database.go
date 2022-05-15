package config

import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
// trigger build 2
type Category struct {
	Id   string
	Name string
}

func (Category) TableName() string {
	return "category"
}

func ConnectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("admin:Admin123@tcp(database-mysql.cw5vouz1a62c.us-west-1.rds.amazonaws.com:3306)/mydb")

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
	err := db.AutoMigrate(model.User{})
	if err != nil {
		return err
	}
	// err = db.Create(model.User{
	// 	Username: "melati",
	// 	Password: "1234",
	// 	Name:     "melati",
	// 	Role:     "admin",
	// }).Error
	return err
}
