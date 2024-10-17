package pokeapi_test

import (
	"bytes"
	"io"
	"net/http"
)

type mockHttpClient struct {
	wasCalledWithUrls []string
	shouldReturn      []byte
	shouldError       error
	returnStatus      int
}

func (c *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	statusCode := 200
	if c.returnStatus != 0 {
		statusCode = c.returnStatus
	}

	c.wasCalledWithUrls = append(c.wasCalledWithUrls, req.URL.String())
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewReader(c.shouldReturn)),
	}, c.shouldError
}
