package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	app := fiber.New()

	app.Test(httptest.NewRequest("GET", "/", nil))
	app.Test(httptest.NewRequest("GET", "/health", nil))
}
