package commands_test

import (
	"bytes"
	"io"
	"net/http"
)

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
