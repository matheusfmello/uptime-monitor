package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusfmello/uptime-monitor/internal/handlers"
	"github.com/matheusfmello/uptime-monitor/internal/middleware"
	"github.com/matheusfmello/uptime-monitor/internal/services"
)

func SetupRoutes(
	app *fiber.App,
	services *services.ServiceFactory,
) {
	authHandler := handlers.NewAuthHandler(services.UserService)
	monitor_handler := handlers.NewMonitorHandler(services.MonitorService)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	monitors := api.Group("/monitors", middleware.JWTMiddleware())
	monitors.Post("", monitor_handler.Create)
	monitors.Get("", monitor_handler.List)

}
