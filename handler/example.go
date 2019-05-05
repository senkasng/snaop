package handler

  
import (
	"github.com/labstack/echo"
	"net/http"
	//"fmt"
)

func Example(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}