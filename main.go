package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// new echo instance
	e := echo.New()

	// to import ent, write the following code at go.mod
	// require github.com/facebook/ent v0.5.0

	e.Use(middleware.CORS())

	// new echo handler for / path
	e.GET("/", func(c echo.Context) error {

		//set allow access control origin header
		c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		// Access-Control-Allow-Methodsand Access-Control-Allow-Headersshould contain the same value
		//as requested in Access-Control-request-Methodsand Access-Control-request-Headersrespectively.
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, *")

		// get whole set-cookie from echo context
		cookies := c.Cookies()
		for _, cookie := range cookies {
			fmt.Println(cookie.Name + cookie.Value)
		}

		return c.NoContent(http.StatusOK)
	})

	// start echo server
	e.Logger.Fatal(e.Start(":9091"))
}
