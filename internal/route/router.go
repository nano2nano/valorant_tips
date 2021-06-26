package route

import (
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
	"github.com/nano2nano/valorant_tips/internal/api"
	"github.com/nano2nano/valorant_tips/internal/db"
	myMw "github.com/nano2nano/valorant_tips/internal/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(echoMw.Logger())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	e.GET("/", api.Status())
	return e
}
