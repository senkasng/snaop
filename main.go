package main

import (
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./routes"
	"./util"
)


func main() {
	e := echo.New()
	//e.Static("/js","./vendor/skyoof/js")
	util.SetLogConfig(e)
	e.Use(middleware.Recover())
	global := util.ParseGlobal()
	//e.File("/","./vendor/index.html")
	//e.File("/favicon.ico", "./vendor/skyoof/favicon.ico")
	for _,route := range routes.Routes {	
		if route.Get != nil {e.GET(route.Url,route.Get)}
		if route.Put != nil {e.PUT(route.Url,route.Put)}
		if route.Post != nil {e.POST(route.Url,route.Post)}
		if route.Delete != nil {e.DELETE(route.Url,route.Delete)}
	}
	e.Logger.Fatal(e.Start(global["ip"]+":"+global["port"]))
}
