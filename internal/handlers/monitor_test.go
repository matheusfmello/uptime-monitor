package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/matheusfmello/uptime-monitor/internal/models"
	"github.com/matheusfmello/uptime-monitor/internal/services"
)

type MockMonitorService struct {
	mock.Mock
}

func (m *MockMonitorService) Create(userID uint, url string, freq int) (*models.Monitor, error) {
	args := m.Called(userID, url, freq)
	return args.Get(0).(*models.Monitor), args.Error(1)
}

func (m *MockMonitorService) ListByUser(userID uint) ([]models.Monitor, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Monitor), args.Error(1)
}

func (m *MockMonitorService) GetDueMonitors() ([]services.MonitorInfo, error) {
	args := m.Called()
	return args.Get(0).([]services.MonitorInfo), args.Error(1)
}

func (m *MockMonitorService) SaveCheckResult(id uint, status string, responseTime int64) error {
	args := m.Called(id, status, responseTime)
	return args.Error(0)
}

func setupMonitorApp(handler *MonitorHandler) *fiber.App {
	app := fiber.New()
	app.Post("/monitors", func(c *fiber.Ctx) error {
		c.Locals("user_id", "1")
		return handler.Create(c)
	})
	app.Get("/monitors", func(c *fiber.Ctx) error {
		c.Locals("user_id", "1")
		return handler.List(c)
	})
	return app
}

func TestMonitorCreate_Success(t *testing.T) {
	mockService := new(MockMonitorService)
	handler := NewMonitorHandler(mockService)
	app := setupMonitorApp(handler)

	input := map[string]interface{}{
		"url":               "https://example.com",
		"frequency-minutes": 5,
	}
	body, _ := json.Marshal(input)

	mockMonitor := &models.Monitor{URL: "https://example.com", FrequencyMinutes: 5}
	mockService.On("Create", uint(1), "https://example.com", 5).Return(mockMonitor, nil)

	req := httptest.NewRequest(http.MethodPost, "/monitors", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestMonitorList_Success(t *testing.T) {
	mockService := new(MockMonitorService)
	handler := NewMonitorHandler(mockService)
	app := setupMonitorApp(handler)

	mockMonitors := []models.Monitor{
		{ID: 1, URL: "https://a.com", FrequencyMinutes: 3},
		{ID: 2, URL: "https://b.com", FrequencyMinutes: 10},
	}
	mockService.On("ListByUser", uint(1)).Return(mockMonitors, nil)

	req := httptest.NewRequest(http.MethodGet, "/monitors", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetDueMonitors(t *testing.T) {
	mockService := new(MockMonitorService)
	expected := []services.MonitorInfo{
		{ID: 1, URL: "https://a.com"},
		{ID: 2, URL: "https://b.com"},
	}
	mockService.On("GetDueMonitors").Return(expected, nil)

	actual, err := mockService.GetDueMonitors()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	mockService.AssertExpectations(t)
}

func TestSaveCheckResult(t *testing.T) {
	mockService := new(MockMonitorService)
	mockService.On("SaveCheckResult", uint(1), "up", int64(120)).Return(nil)

	err := mockService.SaveCheckResult(1, "up", 120)
	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}
