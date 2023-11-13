package main

import (
	"github.com/Sun-FreePort/Cities-and-Citizen/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHelloRoute(t *testing.T) {
	tests := []struct {
		description  string
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			route:        "/ping",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/notfound",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	router.RegisterB2E(app)

	for _, test := range tests {
		// 利用 httptest 包生成 request
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
