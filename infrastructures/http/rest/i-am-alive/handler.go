package i_am_alive

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func IAmAliveHandler(c echo.Context) error {
	return c.String(http.StatusOK, "I AM Alive")
}
