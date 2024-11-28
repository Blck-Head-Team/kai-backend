package mysql

import (
	"fmt"
	"kai-backend/repository/mysql/newsletters"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_USERNAME": viper.GetString("DB_USERNAME"),
		"DB_PASSWORD": viper.GetString("DB_PASSWORD"),
		"DB_HOST":     viper.GetString("DB_HOST"),
		"DB_PORT":     viper.GetString("DB_PORT"),
		"DB_NAME":     viper.GetString("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_USERNAME"], config["DB_PASSWORD"], config["DB_HOST"], config["DB_PORT"], config["DB_NAME"])

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&newsletters.Newsletter{},
	)
	return DB
}
