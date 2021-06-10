package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetUserManagementsController(c echo.Context) error {
	userManagements, err := database.GetUserManagements()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":          "success",
		"userManagements": userManagements,
	})
}

func GetUserManagementController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a userManagement, userManagement with ID " + c.Param("id") + " is not found",
		})
	}

	userManagement, getErr := database.GetUserManagement(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"userManagement": userManagement,
	})
}

func DeleteUserManagementController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a userManagement, userManagement with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeleteUserManagement(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func UpdateUserManagementController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a userManagement, userManagement with ID " + c.Param("id") + " is not found",
		})
	}

	var updateUserManagement models.UserManagement
	c.Bind(&updateUserManagement)
	userManagement, updateErr := database.UpdateUserManagement(id, &updateUserManagement)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"userManagement": userManagement,
	})
}

func CreateUserManagementController(c echo.Context) error {
	userManagement := models.UserManagement{}
	c.Bind(&userManagement)

	_, err := database.CreateUserManagement(&userManagement)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"userManagement": userManagement,
	})
}

func LoginUserManagementsController(c echo.Context) error {
	userManagement := models.UserManagement{}
	c.Bind(&userManagement)

	userManagements, e := database.LoginUserManagements(&userManagement)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":          "success login",
		"userManagements": userManagements,
	})
}
