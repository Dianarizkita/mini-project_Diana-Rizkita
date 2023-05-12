package routes

import (
	"net/http"
	"task/constants"
	"task/controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controllers.LoginUserController)

	//routes Book
	book := e.Group("/books", middleware.JWT([]byte(constants.SECRET_JWT)))
	book.GET("", controllers.GetBookControllers)
	book.POST("", controllers.CreateBookController)
	book.GET(":id", controllers.GetBookController)
	book.PUT(":id", controllers.UpdateBookController)
	book.DELETE(":id", controllers.DeleteBookController)

	//route book details
	book_details := e.Group("/book_details")
	book_details.GET("", controllers.GetBookDetailsControllers)
	book_details.POST("", controllers.CreateBookDetailsController)
	book_details.GET(":id", controllers.GetBookDetailsController)
	book_details.PUT(":id", controllers.UpdateBookDetailsController)
	book_details.DELETE(":id", controllers.DeleteBookDetailsController)

	//route book returns
	book_return := e.Group("/book_return")
	book_return.GET("", controllers.GetReturnControllers)
	book_return.POST("", controllers.CreateReturnController)
	book_return.GET(":id", controllers.GetReturnController)
	book_return.PUT(":id", controllers.UpdateReturnController)
	book_return.DELETE(":id", controllers.DeleteReturnController)

	//route user
	user := e.Group("/users")
	user.GET("", controllers.GetUserControllers)
	user.POST("", controllers.CreateUserController)
	user.GET(":id", controllers.GetUserController)
	user.PUT(":id", controllers.UpdateUserController)
	user.DELETE(":id", controllers.DeleteUserController)

	//route transaction
	transaction := e.Group("/transactions")
	transaction.GET("", controllers.GetTransactionControllers)
	transaction.POST("", controllers.CreateTransactionController)
	transaction.GET(":id", controllers.GetTransactionController)
	transaction.PUT(":id", controllers.UpdateTransactionController)
	transaction.DELETE(":id", controllers.DeleteTransactionController)
	return e
}
