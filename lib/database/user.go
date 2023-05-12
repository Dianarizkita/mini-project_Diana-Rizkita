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

func GetUsers() (users []models.User, err error) {
	if err = config.DB.Model(&models.User{}).Preload("Transactions").Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUser(id uint) (user models.User, err error) {
	user.ID = id
	if err = config.DB.Model(&models.User{}).Preload("Transactions").First(&user).Error; err != nil {
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

func LoginUser(user *models.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
