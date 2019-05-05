package util

import (
	"os"
	//"../config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetLogConfig( e *echo.Echo ) {
	global := ParseGlobal()
	logfile := global["logfile"]
	file,err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic (err)
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
		`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		`"bytes_out":${bytes_out},"referer":${referer},"user_agent":${user_agent},"header":${header}}` + "\n",
		Output: file,
	}))
}