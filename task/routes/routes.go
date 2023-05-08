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

	//route book details
	book_details := e.Group("/book_details")
	book_details.GET("", controllers.GetBookDetailsControllers)
	book_details.GET("all", controllers.GetBookDet)
	book_details.POST("", controllers.CreateBookDetailsController)
	book_details.GET(":id", controllers.GetBookDetailsController)
	book_details.PUT(":id", controllers.UpdateBookDetailsController)
	book_details.DELETE(":id", controllers.DeleteBookDetailsController)

	//route transaction
	transaction := e.Group("/transactions")
	transaction.GET("", controllers.GetTransactionControllers)
	transaction.POST("", controllers.CreateTransactionController)
	transaction.GET(":id", controllers.GetTransactionController)
	transaction.PUT(":id", controllers.UpdateTransactionController)
	transaction.DELETE(":id", controllers.DeleteTransactionController)

	//route user
	user := e.Group("/users")
	user.GET("", controllers.GetUserControllers)
	user.POST("", controllers.CreateUserController)
	user.GET(":id", controllers.GetUserController)
	user.PUT(":id", controllers.UpdateUserController)
	user.DELETE(":id", controllers.DeleteUserController)

	return e
}
