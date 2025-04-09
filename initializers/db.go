package initializers

import (
	"log"
	"os"

	"task-manager/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Faild to connect to DB")
	}

	DB.AutoMigrate(&models.User{}, &models.Task{})
}
