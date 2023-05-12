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

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controllers.LoginUserController)

	//routes Book
	book := e.Group("/books")
	book.GET("/all", controllers.GetBookControllers)
	book.POST("/add", controllers.CreateBookController)
	book.GET(":id", controllers.GetBookController)
	book.PUT("/edit/:id", controllers.UpdateBookController)
	book.DELETE("/delete/:id", controllers.DeleteBookController)

	//route book details
	book_details := e.Group("/book_details")
	book_details.GET("/all", controllers.GetBookDetailsControllers)
	book_details.POST("/add", controllers.CreateBookDetailsController)
	book_details.GET(":id", controllers.GetBookDetailsController)
	book_details.PUT("/edit/:id", controllers.UpdateBookDetailsController)
	book_details.DELETE("/delete/:id", controllers.DeleteBookDetailsController)

	//route book returns
	book_return := e.Group("/book_return")
	book_return.GET("/all", controllers.GetReturnControllers)
	book_return.POST("/add", controllers.CreateReturnController)
	book_return.GET(":id", controllers.GetReturnController)
	book_return.PUT("/edit/:id", controllers.UpdateReturnController)
	book_return.DELETE("/delete/:id", controllers.DeleteReturnController)

	//route user
	user := e.Group("/users")
	user.GET("/all", controllers.GetUserControllers)
	user.POST("/add", controllers.CreateUserController)
	user.GET(":id", controllers.GetUserController)
	user.PUT("/edit/:id", controllers.UpdateUserController)
	user.DELETE("/delete/:id", controllers.DeleteUserController)

	//route transaction
	transaction := e.Group("/transactions")
	transaction.GET("/all", controllers.GetTransactionControllers)
	transaction.POST("/add", controllers.CreateTransactionController)
	transaction.GET(":id", controllers.GetTransactionController)
	transaction.PUT("/edit/:id", controllers.UpdateTransactionController)
	transaction.DELETE("/delete/:id", controllers.DeleteTransactionController)
	return e
}
