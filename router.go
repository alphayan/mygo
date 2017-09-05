package main

import (
	"github.com/echo-contrib/echopprof"
	"github.com/labstack/echo"
)

func initrouter() {
	e := echo.New()
	e.Static("/layui", "layui")
	e.Static("/static", "static")
	//e.Use(middleware.Logger())
	//e.GET("/",func(c echo.Context)error{
	//	return c.Render(200,"index.html",nil)
	//})
	e.File("/", "views/index.html")
	e.File("/a", "views/index1.html")
	e.POST("/login", login)
	echopprof.Wrapper(e)
	e.Start(":8080")
}
