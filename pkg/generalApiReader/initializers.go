package generalApiReader

import (
	"net/http"
	"net/url"
)

func CreateRequest(endpoint string, urlValues, headerValues map[string]string, method string) (*http.Request, error) {
	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return &http.Request{}, err
	}
	params := url.Values{}
	for key, val := range urlValues {
		params.Set(key, val)
	}
	baseURL.RawQuery = params.Encode()

	r, err := http.NewRequest(method, baseURL.String(), nil)
	if err != nil {
		return &http.Request{}, err
	}

	for key, val := range headerValues {
		params.Set(key, val)
	}

	return r, nil
}

func CreateGetRequest(endpoint string, urlValues, headerValues map[string]string) (*http.Request, error) {
	return CreateRequest(endpoint, urlValues, headerValues, "GET")
}

