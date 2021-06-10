package controllers

import (
	"net/http"
	"strconv"

	"project/lib/database"
	"project/models"

	"github.com/labstack/echo"
)

func GetUserAdminsController(c echo.Context) error {
	userAdmins, err := database.GetUserAdmins()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"userAdmins": userAdmins,
	})
}

func GetUserAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a userAdmin, userAdmin with ID " + c.Param("id") + " is not found",
		})
	}

	userAdmin, getErr := database.GetUserAdmin(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"userAdmin": userAdmin,
	})
}

func DeleteUserAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a userAdmin, userAdmin with ID " + c.Param("id") + " is not found",
		})
	}

	if _, deleteErr := database.DeleteUserAdmin(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func UpdateUserAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a userAdmin, userAdmin with ID " + c.Param("id") + " is not found",
		})
	}

	var updateUserAdmin models.UserAdmin
	c.Bind(&updateUserAdmin)
	userAdmin, updateErr := database.UpdateUserAdmin(id, &updateUserAdmin)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"userAdmin": userAdmin,
	})
}

func CreateUserAdminController(c echo.Context) error {
	userAdmin := models.UserAdmin{}
	c.Bind(&userAdmin)

	_, err := database.CreateUserAdmin(&userAdmin)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"userAdmin": userAdmin,
	})
}

func LoginUserAdminsController(c echo.Context) error {
	userAdmin := models.UserAdmin{}
	c.Bind(&userAdmin)

	userAdmins, e := database.LoginUserAdmins(&userAdmin)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success login",
		"userAdmins": userAdmins,
	})
}
