package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetPermissionTransactionsController(c echo.Context) error {
	permissionTransactions, err := database.GetPermissionTransactions()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":                 "success",
		"permissionTransactions": permissionTransactions,
	})
}

func GetPermissionTransactionController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a permissionTransaction, permissionTransaction with ID " + c.Param("id") + " is not found",
		})
	}

	permissionTransaction, getErr := database.GetPermissionTransaction(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":                "success",
		"permissionTransaction": permissionTransaction,
	})
}

func DeletePermissionTransactionController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a permissionTransaction, permissionTransaction with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeletePermissionTransaction(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func UpdatePermissionTransactionController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a permissionTransaction, permissionTransaction with ID " + c.Param("id") + " is not found",
		})
	}

	var updatePermissionTransaction models.PermissionTransaction
	c.Bind(&updatePermissionTransaction)
	permissionTransaction, updateErr := database.UpdatePermissionTransaction(id, &updatePermissionTransaction)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":                "success",
		"permissionTransaction": permissionTransaction,
	})
}

func CreatePermissionTransactionController(c echo.Context) error {
	permissionTransaction := models.PermissionTransaction{}
	c.Bind(&permissionTransaction)

	_, err := database.CreatePermissionTransaction(&permissionTransaction)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":                "success",
		"permissionTransaction": permissionTransaction,
	})
}
