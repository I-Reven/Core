package i_am_alive

import (
	handler "github.com/I-Reven/Core/infrastructures/http/rest/i-am-alive"
	"github.com/labstack/echo/v4"
)

func Boot(e *echo.Echo)  {
	e.GET("/i-am-alive", handler.IAmAliveHandler)
}
