package routes

import (
	"project/constants"
	"project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUsersController)
	e.POST("/users", controllers.CreateUserController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.PUT("/users/add/:id", controllers.AddUserController)
	r.POST("/users", controllers.CreateUserController)

	r.GET("/password_forgets", controllers.GetPasswordForgetsController)
	r.GET("/password_forgets/:id", controllers.GetPasswordForgetController)
	r.DELETE("/password_forgets/:id", controllers.DeletePasswordForgetController)
	r.PUT("/password_forgets/:id", controllers.UpdatePasswordForgetController)
	r.POST("/password_forgets", controllers.CreatePasswordForgetController)

	r.GET("/permission_masters", controllers.GetPermissionMastersController)
	r.GET("/permission_masters/:id", controllers.GetPermissionMasterController)
	r.DELETE("/permission_masters/:id", controllers.DeletePermissionMasterController)
	r.PUT("/permission_masters/:id", controllers.UpdatePermissionMasterController)
	r.POST("/permission_masters", controllers.CreatePermissionMasterController)

	r.GET("/permission_transactions", controllers.GetPermissionTransactionsController)
	r.GET("/permission_transactions/:id", controllers.GetPermissionTransactionController)
	r.DELETE("/permission_transactions/:id", controllers.DeletePermissionTransactionController)
	r.PUT("/permission_transactions/:id", controllers.UpdatePermissionTransactionController)
	r.POST("/permission_transactions", controllers.CreatePermissionTransactionController)

	return e
}
