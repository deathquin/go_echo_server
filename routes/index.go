package routes

import (
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"time"
)

// Handler
func sayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func sayHello2(c echo.Context) error {

	var content struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}

	content.Response = "Sent via JSONP"
	content.Timestamp = time.Now().UTC()
	content.Random = rand.Intn(1000)

	return c.JSON(http.StatusOK, &content)
}

func IndexRoutes(group *echo.Group) {
	/*group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})*/

	group.GET("", sayHello)
	group.GET("/hello", sayHello2)
}
