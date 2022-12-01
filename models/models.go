package models

type EnvVariables struct {
	Debug   bool
	Port    string `default:"8080"`
	BaseURL string `split_words:"true" default:"127.0.0.1:8080"`
}

type URLShortenRequest struct {
	URL string `json:"url"`
}

type URLShortnerResponse struct {
	URL string `json:"url"`
}
