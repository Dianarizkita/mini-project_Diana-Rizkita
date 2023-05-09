package controllers

import (
	"net/http"
	"strconv"
	"task/lib/database"
	"task/models"
	"task/usecase"

	"github.com/labstack/echo/v4"
)

func GetReturnControllers(c echo.Context) error {
	returns, e := database.GetBookReturns()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   returns,
	})
}

func GetReturnController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	returns, err := database.GetBookReturn(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"returns": returns,
	})

}

func CreateReturnController(c echo.Context) error {
	returns := models.Book_Return{}
	c.Bind(&returns)
	if err := usecase.CreateBookReturn(&returns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create book return",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Create New books",
		"returns": returns,
	})
}

func DeleteReturnController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteBookReturn(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error elete book return",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf  pengembalian tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete book returns",
	})
}

func UpdateReturnController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	returns, err := usecase.GetBookReturn(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&returns)

	if err := usecase.UpdateBookReturn(&returns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create pengembalian",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf buku tidak dapat dikembalikan",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"returns": returns,
	})
}
