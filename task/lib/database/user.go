package database

import (
	"task/config"
	"task/models"
)

func CreateUser(user *models.User) error {
	if err := config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id uint) (user models.User, err error) {
	user.ID = id
	if err = config.DB.First(&user).Error; err != nil {
		return
	}
	return
}

//	func UpdateUser(id uint) (user models.User, err error) {
//		user.ID = id
//		if err = config.DB.Updates(&user).Error; err != nil {
//			return
//		}
//		return
//	}
func UpdateUser(user *models.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
