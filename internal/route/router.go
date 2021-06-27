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
	v1 := e.Group("/api/v1")
	agent := v1.Group("/agent")
	agent.GET("", api.GetAgents())
	agent.GET("/:id", api.GetAgent())
	ability := v1.Group("/ability")
	ability.GET("", api.GetAbilities())
	ability.GET("/:id", api.GetAbility())
	_map := v1.Group("/map")
	_map.GET("", api.GetMaps())
	// _map.POST("", api.PostMap())
	_map.GET("/:id", api.GetMap())
	return e
}
