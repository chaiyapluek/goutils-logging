package middlewares

import (
	"io"

	"github.com/chaiyapluek/goutils-logging/src/utils"
	"github.com/labstack/echo/v4"
)

func PrepareRequestContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.NewEchoContext(c)
		return next(ctx)
	}
}

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := utils.GetLogger()
		ctx := c.(*utils.EchoContext)
		LogRequest(log, ctx)
		next(ctx)
		return nil
	}
}

func LogRequest(log *utils.Logger, c *utils.EchoContext) {
	s := c.Request().Method + " " + c.Request().URL.String()
	log.CInfo(c, "Incoming request %s", s)
	if c.Request().Method == "POST" || c.Request().Method == "PUT" {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			tmp, err := utils.MaskAndOmmit(b)
			if err == nil {
				s += "\nrequest body: " + tmp
			}
		}
	}
}
