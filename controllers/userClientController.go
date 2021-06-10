package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetUserClientsController(c echo.Context) error {
	userClients, err := database.GetUserClients()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"userClients": userClients,
	})
}

func GetUserClientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a userClient, userClient with ID " + c.Param("id") + " is not found",
		})
	}

	userClient, getErr := database.GetUserClient(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"userClient": userClient,
	})
}

func DeleteUserClientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a userClient, userClient with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeleteUserClient(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func UpdateUserClientController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a userClient, userClient with ID " + c.Param("id") + " is not found",
		})
	}

	var updateUserClient models.UserClient
	c.Bind(&updateUserClient)
	userClient, updateErr := database.UpdateUserClient(id, &updateUserClient)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"userClient": userClient,
	})
}

func CreateUserClientController(c echo.Context) error {
	userClient := models.UserClient{}
	c.Bind(&userClient)

	_, err := database.CreateUserClient(&userClient)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"userClient": userClient,
	})
}

func LoginUserClientsController(c echo.Context) error {
	userClient := models.UserClient{}
	c.Bind(&userClient)

	userClients, e := database.LoginUserClients(&userClient)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success login",
		"userClients": userClients,
	})
}
