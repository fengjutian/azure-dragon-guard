package repository

import (
	"github.com/fengjutian/azure-dragon-guard/database"
	"github.com/fengjutian/azure-dragon-guard/models"
)

func GetAllUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}

func CreateUser(user *models.User) {
	database.DB.Create(user)
}
