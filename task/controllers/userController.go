package controllers

import (
	"net/http"
	"strconv"
	"task/models"
	"task/usecase"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	users, e := usecase.GetListUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetUser(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"users":   user,
	})

}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := usecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Create New users",
		"users":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error elete user",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf user tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete users",
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	//id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetUser(uint(id))
	//user := models.User{}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create user",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf user tidak dapat diubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
