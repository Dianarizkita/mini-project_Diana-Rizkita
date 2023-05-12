package controllers

import (
	"net/http"
	"strconv"
	"task/models"
	"task/usecase"

	"github.com/labstack/echo/v4"
)

func GetBookControllers(c echo.Context) error {
	books, e := usecase.GetListBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := usecase.GetBook(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"books":   book,
	})

}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := usecase.CreateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Create New books",
		"books":   book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error elete book",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf book tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete books",
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := usecase.GetBook(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&book)

	if err := usecase.UpdateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create book details",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf detail buku tidak dapat diubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
