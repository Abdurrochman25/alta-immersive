package api

import (
	"alta/book-api/api/controllers/book"
	"alta/book-api/api/controllers/customer"
	"alta/book-api/api/controllers/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, customerController *customer.Controller, userController *user.Controller, bookController *book.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/customers/register", customerController.PostCustomerController)
	e.POST("/customers/login", customerController.DeleteCustomerController)
	e.POST("/users/register", userController.PostUserController)
	e.POST("/books/register", bookController.PostBookController)

	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	e.GET("/customers", customerController.GetAllCustomerController)
	e.GET("/customers/:id", customerController.GetCustomerController)
	e.PUT("/customers/:id", customerController.EditCustomerController)
	e.DELETE("/customers/:id", customerController.DeleteCustomerController)

	// ------------------------------------------------------------------
	// CRUD User
	// ------------------------------------------------------------------
	e.GET("/users", userController.GetAllUserController)
	e.GET("/users/:id", userController.GetUserController)
	e.PUT("/users/:id", userController.EditUserController)
	e.DELETE("/users/:id", userController.DeleteUserController)

	// ------------------------------------------------------------------
	// CRUD Book
	// ------------------------------------------------------------------
	e.GET("/books", bookController.GetAllBookController)
	e.GET("/books/:id", bookController.GetBookController)
	e.PUT("/books/:id", bookController.EditBookController)
	e.DELETE("/books/:id", bookController.DeleteBookController)
}
