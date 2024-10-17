package pokeapi

import (
	"errors"
	"net/http"
)

var ErrRequestInvalid = errors.New("request was invalid")
var ErrRequestFailed = errors.New("request failed")

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
