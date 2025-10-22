package database

import (
	"fmt"
	"log"

	"github.com/fengjutian/azure-dragon-guard/config"
	"github.com/fengjutian/azure-dragon-guard/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbConfig := config.GetDBConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Connected to PostgreSQL database")

	// 自动迁移模型
	database.AutoMigrate(&models.User{})

	DB = database
}
