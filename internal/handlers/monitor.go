package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusfmello/uptime-monitor/internal/services"
)

type MonitorHandler struct {
	service services.MonitorServiceInterface
}

func NewMonitorHandler(s services.MonitorServiceInterface) *MonitorHandler {
	return &MonitorHandler{s}
}

func (h *MonitorHandler) Create(c *fiber.Ctx) error {
	type Request struct {
		URL              string `json:"url"`
		FrequencyMinutes int    `json:"frequency-minutes"`
	}
	var body Request

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	userId := c.Locals("user_id")
	uid, err := strconv.ParseUint(userId.(string), 10, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("User ID %s is invalid", userId.(string)),
		})
	}

	monitor, err := h.service.Create(uint(uid), body.URL, body.FrequencyMinutes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(monitor)

}

func (h *MonitorHandler) List(c *fiber.Ctx) error {

	userId := c.Locals("user_id")
	uid, err := strconv.ParseUint(userId.(string), 10, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	monitors, err := h.service.ListByUser(uint(uid))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(monitors)
}
