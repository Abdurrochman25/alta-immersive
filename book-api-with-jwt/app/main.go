package main

import (
	"alta/book-api/api"
	bookController "alta/book-api/api/controllers/book"
	customerController "alta/book-api/api/controllers/customer"
	userController "alta/book-api/api/controllers/user"
	"alta/book-api/config"
	"alta/book-api/models"
	"alta/book-api/util"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDatabaseConnection(config)

	//initiate user model
	customerModel := models.NewCustomerModel(db)
	userModel := models.NewUserModel(db)
	bookModel := models.NewBookModel(db)

	//initiate user controller
	newCustomerController := customerController.NewController(customerModel)
	newUserController := userController.NewController(userModel)
	newBookController := bookController.NewController(bookModel)

	//create echo http
	e := echo.New()

	//register API path and controller
	api.RegisterPath(e, newCustomerController, newUserController, newBookController)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
