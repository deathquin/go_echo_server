package main

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	api "go_echo_server/routes"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// auto reload
	// https://devcheat.tistory.com/7
	// Install - go get -u github.com/cosmtrek/air
	// run command - air

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/asset", "public")

	// api routes
	apiGroup := e.Group("/api")
	api.IndexRoutes(apiGroup)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// Named route "foobar"
	e.GET("/singlefile", func(c echo.Context) error {
		return c.Render(http.StatusOK, "singlefile.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	// Named route "foobar"
	e.GET("/multiplefile", func(c echo.Context) error {
		return c.Render(http.StatusOK, "multiplefile.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.Logger.Fatal(e.Start(":1323")) // localhost:1323

}
