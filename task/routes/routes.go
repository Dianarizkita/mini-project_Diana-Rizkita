package routes

import (
	"net/http"
	"task/controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() *echo.Echo {
	e := echo.New()
	//routes Book
	book := e.Group("/books")
	book.GET("", controllers.GetBookControllers)
	book.POST("", controllers.CreateBookController)
	book.GET(":id", controllers.GetBookController)
	book.PUT(":id", controllers.UpdateBookController)
	book.DELETE(":id", controllers.DeleteBookController)

	return e
}
