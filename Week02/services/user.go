package services

import (
	"Go-000/Week02/dao"
	"Go-000/Week02/models"
)

func GetUser(id int) (*models.User, error) {
	return dao.GetUserById(id)
}
