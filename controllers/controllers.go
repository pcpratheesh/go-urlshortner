package controllers

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/pcpratheesh/go-urlshortner/models"
	"github.com/pcpratheesh/go-urlshortner/pkg/shortner"
	"github.com/pcpratheesh/go-urlshortner/pkg/storage"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	baseURL string
	store   storage.StorageInterface
}

type Config struct {
	BaseURL, Store string
}

// NewController - Initialize new controller
func NewController(cfg Config) (*controller, error) {
	// choose the store
	store, err := storage.NewStorage(cfg.Store)
	if err != nil {
		return nil, err
	}

	return &controller{
		cfg.BaseURL, store,
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
// @Produce json
// @Param RequestBody body models.URLShortenRequest true "The body to request an encode"
// @Success 200 {object} models.URLShortenResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 404 {object} models.HTTPError404
// @Failure 500 {object} models.HTTPError500
// @Router /encode [POST]
func (c *controller) EncodeURL(ctx *fiber.Ctx) error {
	logrus.Info("[ENDPOINT] EncodeURL")

	// check for the incoming request body
	body := new(models.URLShortenRequest)
	if err := ctx.BodyParser(&body); err != nil {
		logrus.Errorf("cannot parse request, %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request",
		})
	}

	// check the url is valid
	if !govalidator.IsURL(body.URL) {
		logrus.Errorf("Invalid URL %v", body.URL)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	// check the url is already generated one,
	// then return the existing one without generate new one
	if url, ok := c.store.CheckURLExists(body.URL); ok {
		return ctx.JSON(models.URLShortenResponse{
			URL:  fmt.Sprintf("%s/short/%s", c.baseURL, url),
			Code: url,
		})
	}

	// Generate the shorten url
	generatedURL := shortner.GenerateShortLink(body.URL)

	// store the data
	err := c.store.SaveURL(generatedURL, body.URL)
	if err != nil {
		logrus.Errorf("unable to store the generated url %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unable to store the generated url",
		})
	}

	return ctx.JSON(models.URLShortenResponse{
		URL:  fmt.Sprintf("%s/short/%s", c.baseURL, generatedURL),
		Code: generatedURL,
	})
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
	shortURL := ctx.Params("url")
	logrus.Info("[ENDPOINT] DecodeURL", logrus.WithField("url", shortURL))

	// check the url exists
	originalURl := c.store.RetrieveURL(shortURL)
	if originalURl == "" {
		logrus.Errorf("invalid url %v", shortURL)
		return ctx.Status(fiber.StatusNotFound).SendString("invalid url")
	}

	return ctx.Redirect(originalURl)
}
