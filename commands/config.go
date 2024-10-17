package commands

import (
	"io"
	"net/http"
	"os"

	"github.com/yeldiRium/learning-go-pokedex/model"
	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type CliConfig struct {
	output     io.Writer
	httpClient HttpClient
	cache      pokecache.Cache
	pokedex    map[string]model.Pokemon

	nextMapUrl     *string
	previousMapUrl *string
}

func NewCliConfig() *CliConfig {
	nextUrl := pokeapi.BaseAreaListUrl
	return &CliConfig{
		output:     os.Stdout,
		httpClient: http.DefaultClient,
		pokedex:    make(map[string]model.Pokemon),

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
