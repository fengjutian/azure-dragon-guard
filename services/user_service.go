package services

import "github.com/fengjutian/azure-dragon-guard/models"

var userRepo = []models.User{
	{ID: 1, Name: "张三", Email: "zhangsan@example.com"},
	{ID: 2, Name: "李四", Email: "lisi@example.com"},
	{ID: 3, Name: "王五", Email: "wangwu@example.com"},
}

type User = models.User

func GetAllUsers() []User {
	return userRepo
}

func CreateUser(u *User) {
	userRepo = append(userRepo, *u)
}
