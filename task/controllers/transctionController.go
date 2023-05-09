package controllers

import (
	"net/http"
	"strconv"
	"task/lib/database"
	"task/models"
	"task/usecase"

	"github.com/labstack/echo/v4"
)

func GetTransactionControllers(c echo.Context) error {
	transactions, e := database.GetTransactions()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get all books",
		"transactions": transactions,
	})
}

func GetTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	transaction, err := database.GetTransaction(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "succes",
		"transactions": transaction,
	})

}

func CreateTransactionController(c echo.Context) error {
	transaction := models.Transaction{}
	c.Bind(&transaction)
	if err := usecase.CreateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success Create New books",
		"transactions": transaction,
	})
}

func DeleteTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteTransaction(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error delete transaction",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf transaction tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete transactions",
	})
}

func UpdateTransactionController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := usecase.GetTransaction(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&transaction)

	if err := usecase.UpdateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create transaction",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf buku tidak dapat diubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success update book",
		"transaction": transaction,
	})
}
