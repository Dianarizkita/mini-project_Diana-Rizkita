package main

import (
	"task/config"
	m "task/middleware"
	"task/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.Logmiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
