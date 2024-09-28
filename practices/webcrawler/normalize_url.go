package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {

	url, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	combinedURL := url.Host + url.Path
	loweredURL := strings.ToLower(combinedURL)
	sanitizedURL := strings.TrimSuffix(loweredURL, "/")

	return sanitizedURL, nil

}
