package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pcpratheesh/go-urlshortner/config"
	"github.com/pcpratheesh/go-urlshortner/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func SetupBase() (app *fiber.App, controller *controller) {
	app = fiber.New()

	APP_NAME := "URLSRT"
	// load the configurations
	cfg, err := config.LoadConfigurations(APP_NAME)
	if err != nil {
		logrus.Fatal(err)
	}

	// initialize the controller
	controller, err = NewController(Config{
		BaseURL: cfg.BaseURL,
		Store:   cfg.Store,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	return app, controller
}

func TestGetEndpoint(t *testing.T) {

	app, controller := SetupBase()
	var tests = []struct {
		name           string
		method         string
		endpoint       string
		requestPayload map[string]string
		handlerFunc    func(ctx *fiber.Ctx) error
		expectedBody   string
		description    string
		expectedCode   int
	}{
		{
			name:           "home-endpoint",
			method:         http.MethodGet,
			endpoint:       "/",
			handlerFunc:    controller.Home,
			requestPayload: nil,
			expectedCode:   http.StatusOK,
			description:    "",
		},
		{
			name:           "health-endpoint",
			method:         http.MethodGet,
			endpoint:       "/health",
			handlerFunc:    controller.Health,
			requestPayload: nil,
			expectedCode:   http.StatusOK,
			description:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			app.Get(tc.endpoint, tc.handlerFunc)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest(tc.method, tc.endpoint, nil)

			// Perform the request plain with the app,
			// the second argument is a request latency
			// (set to -1 for no latency)
			resp, _ := app.Test(req, 1)

			// Verify, if the status code is as expected
			assert.Equalf(t, tc.expectedCode, resp.StatusCode, tc.description)

		})
	}
}

func TestURLEncoding(t *testing.T) {

	app, controller := SetupBase()

	var tests = []struct {
		name           string
		method         string
		endpoint       string
		requestPayload map[string]string
		headers        map[string]string
		handlerFunc    func(ctx *fiber.Ctx) error
		expectedBody   string
		description    string
		expectedCode   int
	}{
		// invalid payload request
		{
			name:        "invalid-payload",
			method:      http.MethodPost,
			endpoint:    "/encode",
			handlerFunc: controller.EncodeURL,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			requestPayload: map[string]string{
				"url": "abc.co;",
			},
			expectedCode: http.StatusBadRequest,
			description:  "",
		},
		// invalid url
		{
			name:        "invalid-url",
			method:      http.MethodPost,
			endpoint:    "/encode",
			handlerFunc: controller.EncodeURL,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			requestPayload: map[string]string{
				"url": "abcdefg",
			},
			expectedCode: http.StatusBadRequest,
			description:  "",
		},
		// invalid url
		{
			name:        "valid-url",
			method:      http.MethodPost,
			endpoint:    "/encode",
			handlerFunc: controller.EncodeURL,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			requestPayload: map[string]string{
				"url": "www.google.com",
			},
			expectedCode: http.StatusOK,
			description:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			app.Post(tc.endpoint, tc.handlerFunc)

			payload, err := json.Marshal(tc.requestPayload)
			assert.Nil(t, err)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest(tc.method, tc.endpoint, bytes.NewBuffer(payload))

			for key, val := range tc.headers {
				req.Header.Add(key, val)
			}

			// Perform the request plain with the app,
			// the second argument is a request latency
			// (set to -1 for no latency)
			resp, _ := app.Test(req, 1)

			// Verify, if the status code is as expected
			assert.Equalf(t, tc.expectedCode, resp.StatusCode, tc.description)

		})
	}
}

func TestCheckingURL(t *testing.T) {
	app, controller := SetupBase()
	var responseModel models.URLShortnerResponse

	t.Run("first-run", func(t *testing.T) {
		app.Post("/encode", controller.EncodeURL)

		payload, err := json.Marshal(map[string]string{
			"url": "www.google.com",
		})
		assert.Nil(t, err)

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(payload))
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)
		defer resp.Body.Close()

		// Verify, if the status code is as expected
		assert.Equalf(t, http.StatusOK, resp.StatusCode, "test")

		// check the data response body
		json.NewDecoder(resp.Body).Decode(&responseModel)
	})

	t.Run("second-run", func(t *testing.T) {
		app.Post("/encode", controller.EncodeURL)

		payload, err := json.Marshal(map[string]string{
			"url": "www.google.com",
		})
		assert.Nil(t, err)

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer(payload))
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)
		defer resp.Body.Close()

		var responseModelSecond models.URLShortnerResponse

		// Verify, if the status code is as expected
		assert.Equalf(t, http.StatusOK, resp.StatusCode, "test")

		// check the data response body
		json.NewDecoder(resp.Body).Decode(&responseModelSecond)

		t.Run("checking-same-code", func(t *testing.T) {
			assert.Equal(t, responseModel.Code, responseModelSecond.Code)
		})

	})
}
