package api

import (
	"alta/book-api/api/controllers/book"
	"alta/book-api/api/controllers/customer"
	"alta/book-api/api/controllers/user"
	"alta/book-api/constants"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, customerController *customer.Controller, userController *user.Controller, bookController *book.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/customers/", customerController.PostCustomerController)
	e.POST("/customers", customerController.PostCustomerController)
	e.POST("/customers/login/", customerController.LoginCustomerController)
	e.POST("/customers/login", customerController.LoginCustomerController)
	e.POST("/users/", userController.PostUserController)
	e.POST("/users", userController.PostUserController)
	e.POST("/users/login/", userController.LoginUserController)
	e.POST("/users/login", userController.LoginUserController)
	e.POST("/books/", bookController.PostBookController)
	e.POST("/books", bookController.PostBookController)

	// ------------------------------------------------------------------
	// JWT Auth
	// ------------------------------------------------------------------
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	eAuth.GET("/customers", customerController.GetAllCustomerController)
	eAuth.GET("/customers/", customerController.GetAllCustomerController)
	eAuth.GET("/customers/:id", customerController.GetCustomerController)
	eAuth.GET("/customers/:id/", customerController.GetCustomerController)
	eAuth.PUT("/customers/:id", customerController.EditCustomerController)
	eAuth.PUT("/customers/:id/", customerController.EditCustomerController)
	eAuth.DELETE("/customers/:id", customerController.DeleteCustomerController)
	eAuth.DELETE("/customers/:id/", customerController.DeleteCustomerController)

	// ------------------------------------------------------------------
	// CRUD User
	// ------------------------------------------------------------------
	eAuth.GET("/users", userController.GetAllUserController)
	eAuth.GET("/users/", userController.GetAllUserController)
	eAuth.GET("/users/:id", userController.GetUserController)
	eAuth.GET("/users/:id/", userController.GetUserController)
	eAuth.PUT("/users/:id", userController.EditUserController)
	eAuth.PUT("/users/:id/", userController.EditUserController)
	eAuth.DELETE("/users/:id", userController.DeleteUserController)
	eAuth.DELETE("/users/:id/", userController.DeleteUserController)

	// ------------------------------------------------------------------
	// CRUD Book
	// ------------------------------------------------------------------
	e.GET("/books", bookController.GetAllBookController)
	e.GET("/books/", bookController.GetAllBookController)
	e.GET("/books/:id", bookController.GetBookController)
	e.GET("/books/:id/", bookController.GetBookController)
	eAuth.PUT("/books/:id", bookController.EditBookController)
	eAuth.PUT("/books/:id/", bookController.EditBookController)
	eAuth.DELETE("/books/:id", bookController.DeleteBookController)
	eAuth.DELETE("/books/:id/", bookController.DeleteBookController)
}
