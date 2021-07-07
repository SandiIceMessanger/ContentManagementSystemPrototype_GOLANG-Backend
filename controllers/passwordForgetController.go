package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetPasswordForgetsController(c echo.Context) error {
	passwordForgets, err := database.GetPasswordForgets()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":          "success",
		"passwordForgets": passwordForgets,
	})
}

func GetPasswordForgetController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a passwordForget, passwordForget with ID " + c.Param("id") + " is not found",
		})
	}

	passwordForget, getErr := database.GetPasswordForget(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"passwordForget": passwordForget,
	})
}

func DeletePasswordForgetController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a passwordForget, passwordForget with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeletePasswordForget(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func UpdatePasswordForgetController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a passwordForget, passwordForget with ID " + c.Param("id") + " is not found",
		})
	}

	var updatePasswordForget models.PasswordForget
	c.Bind(&updatePasswordForget)
	passwordForget, updateErr := database.UpdatePasswordForget(id, &updatePasswordForget)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"passwordForget": passwordForget,
	})
}

func CreatePasswordForgetController(c echo.Context) error {
	passwordForget := models.PasswordForget{}
	c.Bind(&passwordForget)

	_, err := database.CreatePasswordForget(&passwordForget)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":         "success",
		"passwordForget": passwordForget,
	})
}
