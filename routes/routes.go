package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUserManagementsController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	r.GET("/user_managements", controllers.GetUserManagementsController)
	r.GET("/user_managements/:id", controllers.GetUserManagementController)
	r.DELETE("/user_managements/:id", controllers.DeleteUserManagementController)
	r.PUT("/user_managements/:id", controllers.UpdateUserManagementController)
	r.POST("/user_managements", controllers.CreateUserManagementController)

	return e
}
