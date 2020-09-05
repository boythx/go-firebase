package routes

import (
	crud "firebase/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return crud.Home(c)
	})
	e.POST("/", func(c echo.Context) error {
		return crud.AddData(c)
	})
	e.DELETE("/:_id", func(c echo.Context) error {
		return crud.AddData(c)
	})
	e.Logger.Fatal(e.Start(":9000"))
}
