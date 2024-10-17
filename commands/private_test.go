package commands

import (
	"bytes"
	"io"
	"net/http"
)

func (c *CliConfig) WithNextMapUrl(nextUrl string) *CliConfig {
	c.nextMapUrl = &nextUrl
	return c
}

func (c *CliConfig) WithPreviousMapUrl(previousUrl string) *CliConfig {
	c.previousMapUrl = &previousUrl
	return c
}

type mockHttpClient struct {
	wasCalledWithUrls []string
	shouldReturn      []byte
	shouldError       error
}

func (c *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.wasCalledWithUrls = append(c.wasCalledWithUrls, req.URL.String())
	return &http.Response{
		Body: io.NopCloser(bytes.NewReader(c.shouldReturn)),
	}, c.shouldError
}
