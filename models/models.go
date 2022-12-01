package models

type EnvVariables struct {
	Debug   bool
	Port    string `default:"8080"`
	BaseURL string `split_words:"true" default:"https://127.0.0.1:8080"`
	Store   string `split_words:"true" default:"InMemory"`
	XApiKey string `split_words:"true" default:"3cbc5291f1e04ebe5ea24bfdba6763c49c597cea"`
}

type URLShortenRequest struct {
	URL string `json:"url"`
}

type URLShortenResponse struct {
	URL  string `json:"url"`
	Code string `json:"code"`
}
