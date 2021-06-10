package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	r.GET("/user_admins", controllers.GetUserAdminsController)
	r.GET("/user_admins/:id", controllers.GetUserAdminController)
	r.DELETE("/user_admins/:id", controllers.DeleteUserAdminController)
	r.PUT("/user_admins/:id", controllers.UpdateUserAdminController)
	r.POST("/user_admins", controllers.CreateUserAdminController)

	r.GET("/user_clients", controllers.GetUserClientsController)
	r.GET("/user_clients/:id", controllers.GetUserClientController)
	r.DELETE("/user_clients/:id", controllers.DeleteUserClientController)
	r.PUT("/user_clients/:id", controllers.UpdateUserClientController)
	r.POST("/user_clients", controllers.CreateUserClientController)

	r.GET("/user_managements", controllers.GetUserManagementsController)
	r.GET("/user_managements/:id", controllers.GetUserManagementController)
	r.DELETE("/user_managements/:id", controllers.DeleteUserManagementController)
	r.PUT("/user_managements/:id", controllers.UpdateUserManagementController)
	r.POST("/user_managements", controllers.CreateUserManagementController)

	return e
}
