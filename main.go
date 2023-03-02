package main

import "github.com/labstack/echo/v4"

func main() {
	// new echo instance
	e := echo.New()

	// new echo handler for / path
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})

	// start echo server
	e.Logger.Fatal(e.Start(":8080"))
}
