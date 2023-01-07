package router

import "github.com/labstack/echo/v4"

func Route() {
	e := echo.New()
	Articles(e)
	e.Logger.Fatal(e.Start(":1323"))
}
