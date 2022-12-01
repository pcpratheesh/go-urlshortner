package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	baseURL string
}

type Config struct {
	BaseURL string
}

// NewController - Initialize new controller
func NewController(cfg Config) (*controller, error) {
	return &controller{
		cfg.BaseURL,
	}, nil
}

// Home.
// @Description Home endpoint.
// @Summary Home endpoint
// @Success 200 {string} status "ok"
// @Success 200 {object} string
// @Failure 400 {object} models.HTTPError400
// @Failure 404 {object} models.HTTPError404
// @Failure 500 {object} models.HTTPError500
// @Router / [get]
func (c *controller) Home(ctx *fiber.Ctx) error {
	return ctx.SendString("server is running")
}

// HealthCheck godoc
// @Summary Health check endpoint for urlshortner
// @Description This endpoint exists solely for checking the active status of the application. Any HTTP status other than 200 signifies that the application is down
// @Accept */*
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} models.HTTPError400
// @Failure 404 {object} models.HTTPError404
// @Failure 500 {object} models.HTTPError500
// @Router /health [GET]
func (c *controller) Health(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"status": "ok"})
}

// EncodeURL godoc
// @Summary Endpoint encode url
// @Description This endpoint is for prepare teh encoded url
// @Accept */*
// @Produce json
// @Success 200 {object} models.URLShortnerResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 404 {object} models.HTTPError404
// @Failure 500 {object} models.HTTPError500
// @Router /encode [POST]
func (c *controller) EncodeURL(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).SendString("endpoint not implemented")
}

// DecodeURL godoc
// @Summary Endpoint decode url
// @Description fetch the original url and redirect to that web site
// @Accept */*
// @Produce json
// @Param shorturl path string true "shorturl"
// @Success 200 {object} string
// @Failure 400 {object} models.HTTPError400
// @Failure 404 {object} models.HTTPError404
// @Failure 500 {object} models.HTTPError500
// @Router /short/{shorturl} [GET]
func (c *controller) DecodeURL(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).SendString("endpoint not implemented")
}
