package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	user, getErr := database.GetUser(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	var deleteUser models.User
	c.Bind(&deleteUser)
	user, deleteErr := database.DeleteUser(id)
	if deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	var updateUser models.User
	c.Bind(&updateUser)
	user, updateErr := database.UpdateUser(id, &updateUser)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func AddUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	var addUser models.User
	c.Bind(&addUser)
	user, addErr := database.DeleteUser(id)
	if addErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, addErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	_, err := database.CreateUser(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"full name":   user.Name,
		"token":       user.Token,
		"user_status": user.Status,
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
