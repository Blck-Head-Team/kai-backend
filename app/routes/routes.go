package routes

import (
	"kai-backend/app/middlewares"
	"kai-backend/controllers/newsletters"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllersList struct {
	JWTMiddleware         middleware.JWTConfig
	NewslettersController *newsletters.Controllers
}

func (controllers ControllersList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	v1 := e.Group("/api/v1")
	v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	{
		v1.GET("/newsletters", controllers.NewslettersController.GetAll)
		v1.GET("/newsletters/:Id", controllers.NewslettersController.GetById)
		v1.GET("/newsletters/count", controllers.NewslettersController.CountAll)
	}

	admin := v1.Group("", middleware.JWTWithConfig(controllers.JWTMiddleware))
	{
		admin.POST("/newsletters", controllers.NewslettersController.Create, middlewares.OperationalAdmin())
		admin.PUT("/newsletters/:Id", controllers.NewslettersController.Update, middlewares.OperationalAdmin())
		admin.DELETE("/newsletters/:Id", controllers.NewslettersController.Delete, middlewares.OperationalAdmin())
	}
}
