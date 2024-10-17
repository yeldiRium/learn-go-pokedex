package commands

import (
	"io"
	"net/http"
	"os"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type CliConfig struct {
	output     io.Writer
	httpClient HttpClient

	nextMapUrl     *string
	previousMapUrl *string
}

func NewCliConfig() *CliConfig {
	return &CliConfig{
		output:     os.Stdout,
		httpClient: http.DefaultClient,
	}
}

func (c *CliConfig) WithHttpClient(httpClient HttpClient) {
	c.httpClient = httpClient
}
