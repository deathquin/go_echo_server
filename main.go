package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	api "go_echo_server/routes"
	"net/http"
)

func main() {

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// api routes
	apiGroup := e.Group("/api")
	api.IndexRoutes(apiGroup)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323")) // localhost:1323

}
