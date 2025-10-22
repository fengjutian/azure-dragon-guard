package services

import (
	"github.com/fengjutian/azure-dragon-guard/models"
	"github.com/fengjutian/azure-dragon-guard/repository"
)

func GetAllUsers() []models.User {
	return repository.GetAllUsers()
}

func CreateUser(u *models.User) {
	repository.CreateUser(u)
}
