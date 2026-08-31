package assert

import (
	"net/http"
	"net/url"
)

func httpCode(handler http.HandlerFunc, method, url string, values url.Values) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func HTTPSuccess(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPRedirect(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPError(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPStatusCode(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) string {
	_ = "STUB: not implemented"
	return ""
}

func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}
