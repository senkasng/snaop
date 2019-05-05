package routes

import (
	"../handler"
	"github.com/labstack/echo"
)

type route struct {
	Url string
	Get   echo.HandlerFunc
	Put   echo.HandlerFunc
	Post  echo.HandlerFunc
	Delete echo.HandlerFunc
}

var Routes []route

func init()  {
	app := route {
		"/",
		handler.Example,
		nil,
		nil,
		nil,
	}
	Routes = append(Routes,app)
}



