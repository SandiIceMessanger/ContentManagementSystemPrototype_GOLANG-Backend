package main

import (
	"project/config"
	"project/routes"

	m "project/middlewares"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}


