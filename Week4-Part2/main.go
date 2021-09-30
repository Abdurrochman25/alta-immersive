package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" from:"password"`
}

var users []User

// --------------------controller-----------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if id < len(users) && id > -1 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success get user by id",
			"id":       users[id].Id,
			"name":     users[id].Name,
			"email":    users[id].Email,
			"password": users[id].Password,
		})
	}
	return c.String(http.StatusNotFound, "Not Found")
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if id < len(users) && id > -1 {
		for i := id; i < len(users)-1; i++ {
			users[i] = users[i+1]
		}
		users = users[:len(users)-1]
		return c.String(http.StatusOK, "success delete user by id")
	}
	return c.String(http.StatusNotFound, "Not Found")
}

// update user by id
func UpdateUserController(c echo.Context) error {
	ids := c.Param("id")
	user := User{}
	id, _ := strconv.Atoi(ids)
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if id < len(users) && id > -1 {
		users[id] = user
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success update user by id",
			"id":       users[id].Id,
			"name":     users[id].Name,
			"email":    users[id].Email,
			"password": users[id].Password,
		})
	}
	return c.String(http.StatusNotFound, "Not Found")
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

// ---------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
