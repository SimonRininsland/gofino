package main

import (
	"github.com/SimonRininsland/gofino/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// API Path
	v0 := e.Group("/api/v0/")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	//v0.POST("names/sort", handler.SortUser)
	v0.POST("names/add", handler.AddUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
