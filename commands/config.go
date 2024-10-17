package commands

import (
	"io"
	"net/http"
	"os"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
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
	nextUrl := pokeapi.BaseAreaListUrl
	return &CliConfig{
		output:     os.Stdout,
		httpClient: http.DefaultClient,
		nextMapUrl: &nextUrl,
	}
}

func (c *CliConfig) WithHttpClient(httpClient HttpClient) *CliConfig {
	c.httpClient = httpClient
	return c
}

func (c *CliConfig) WithOutput(output io.Writer) *CliConfig {
	c.output = output
	return c
}
