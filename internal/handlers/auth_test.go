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

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(name string, email string, password string) (*models.User, error) {
	args := m.Called(name, email, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) Authenticate(email string, password string) (*models.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func setupApp(userService services.UserServiceInterface) *fiber.App {
	app := fiber.New()
	handler := NewAuthHandler(userService)

	app.Post("/api/auth/register", handler.Register)
	app.Post("/api/auth/login", handler.Login)

	return app
}

func TestRegister(t *testing.T) {
	mockUserService := new(MockUserService)
	mockApp := setupApp(mockUserService)

	mockUser := &models.User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	mockUserService.On("Register", "Alice", "alice@example.com", "password123").Return(mockUser, nil)
	payload := map[string]string{
		"name":     "Alice",
		"email":    "alice@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := mockApp.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
	mockUserService.AssertExpectations(t)

}

func TestLogin(t *testing.T) {
	mockService := new(MockUserService)
	mockApp := setupApp(mockService)

	mockUser := &models.User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}

	mockService.On("Authenticate", "alice@example.com", "password123").Return(mockUser, nil)

	payload := map[string]string{
		"email":    "alice@example.com",
		"password": "password123",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := mockApp.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
	mockService.AssertExpectations(t)
}
