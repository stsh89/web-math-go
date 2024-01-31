package notion

import "log/slog"

type Configuration struct {
	ApiKey     string
	DatabaseId string
}

type Client struct {
	Configuration Configuration
	Logger        *slog.Logger
}
