package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetPermissionMastersController(c echo.Context) error {
	permissionMasters, err := database.GetPermissionMasters()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":            "success",
		"permissionMasters": permissionMasters,
	})
}

func GetPermissionMasterController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a permissionMaster, permissionMaster with ID " + c.Param("id") + " is not found",
		})
	}

	permissionMaster, getErr := database.GetPermissionMaster(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":           "success",
		"permissionMaster": permissionMaster,
	})
}

func DeletePermissionMasterController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a permissionMaster, permissionMaster with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeletePermissionMaster(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func UpdatePermissionMasterController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a permissionMaster, permissionMaster with ID " + c.Param("id") + " is not found",
		})
	}

	var updatePermissionMaster models.PermissionMaster
	c.Bind(&updatePermissionMaster)
	permissionMaster, updateErr := database.UpdatePermissionMaster(id, &updatePermissionMaster)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":           "success",
		"permissionMaster": permissionMaster,
	})
}

func CreatePermissionMasterController(c echo.Context) error {
	permissionMaster := models.PermissionMaster{}
	c.Bind(&permissionMaster)

	_, err := database.CreatePermissionMaster(&permissionMaster)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":           "success",
		"permissionMaster": permissionMaster,
	})
}
