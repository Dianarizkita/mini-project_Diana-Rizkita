package controllers

import (
	"net/http"
	"strconv"
	"task/models"
	"task/usecase"

	"github.com/labstack/echo/v4"
)

func GetBookDetailsControllers(c echo.Context) error {
	books_details, e := usecase.GetListBooksDetails()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success get all books details",
		"books_details": books_details,
	})
}

func GetBookDetailsController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book_details, err := usecase.GetBookDetails(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "succes",
		"books_details": book_details,
	})

}

func CreateBookDetailsController(c echo.Context) error {
	book_details := models.Book_Details{}
	c.Bind(&book_details)
	if err := usecase.CreateBookDetails(&book_details); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create book details",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success Create New books detail",
		"book_details": book_details,
	})
}

func DeleteBookDetailsController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteBookDetails(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error delete book details",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf book details tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete book details",
	})
}

func UpdateBookDetailsController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book_details, err := usecase.GetBookDetails(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&book_details)

	if err := usecase.UpdateBookDetails(&book_details); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create book details",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf detail buku tidak dapat diubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book detail",
		"book":    book_details,
	})
}
