package handler

import (
	"net/http"
	"net/url"
)

type ProxyHandler struct{}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}
}

func (P ProxyHandler) NewProxyClient(rawURL string) (*http.Client, error) {
	proxyURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	return httpClient, nil
}
