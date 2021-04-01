package golinks

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

const (
	baseURLPath   = "/v1"
	testAccessKey = "golinks-test"
)

func setup() (*Client, *http.ServeMux, string, func()) {
	mux := http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	server := httptest.NewServer(apiHandler)
	client := NewClient(testAccessKey)
	url, _ := url.Parse(server.URL + baseURLPath)
	client.BaseURL = url
	return client, mux, server.URL, server.Close
}
