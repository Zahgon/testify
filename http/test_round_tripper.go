package http

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type TestRoundTripper struct {
	mock.Mock
}

func (t *TestRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
