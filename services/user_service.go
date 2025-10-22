package services

import "github.com/fengjutian/azure-dragon-guard/models"

var userRepo = []models.User{}

type User = models.User

func GetAllUsers() []User {
	return userRepo
}

func CreateUser(u *User) {
	userRepo = append(userRepo, *u)
}
