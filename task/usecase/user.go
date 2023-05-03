package usecase

import (
	"errors"
	"fmt"
	"task/lib/database"
	"task/models"
)

func CreateUser(user *models.User) error {
	// check name cannot be empty
	if user.Name == "" {
		return errors.New("user name cannot be empty")
	}

	//check email
	if user.Email == "" {
		return errors.New("user Email cannot be empty")
	}

	err := database.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uint) (user models.User, err error) {
	user.ID = id
	err = database.GetUser(&user)
	if err != nil {
		fmt.Println("Error getting user from database")
		return
	}
	return
}

func GetListUsers() (users []models.User, err error) {
	if err != nil {
		fmt.Println("GetListUser : Error getting user from database")
		return
	}
	return
}

func UpdateUser(user *models.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error Updating User")
		return
	}

	return
}

func DeleteUser(id uint) (err error) {
	user := models.User{}
	user.ID = id
	err = database.DeleteUser(&user)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err ", err)
		return
	}

	return
}

func LoginUser(user *models.User) (token string, err error) {
	//check to DB email and Password
	err = database.GetUser(user)
	if err != nil {
		fmt.Println("Error getting user from database")
		return
	}

	//Generate Jwt

	if err != nil {
		fmt.Println("Error generate token")
		return
	}
	return
}
