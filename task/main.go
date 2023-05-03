package main

import (
	"task/config"
	//m "task/middlewares"
	"task/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	//m.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":8000"))
}
