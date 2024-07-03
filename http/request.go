package http

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpClient struct {
	client *http.Client
}

func New(client *http.Client) *HttpClient {
	return &HttpClient{client}
}

func (c *HttpClient) DoRequest(method, url, body, queryParams, headers string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	err = c.addQueryParams(request, queryParams)
	if err != nil {
		return nil, err
	}

	err = c.addHeaders(request, headers)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *HttpClient) addHeaders(request *http.Request, headers string) error {
	headers = strings.TrimSpace(headers)
	if headers == "" {
		return nil
	}
	for _, line := range strings.Split(headers, "\n") {
		header := strings.Split(line, "=")
		if len(header) != 2 {
			return fmt.Errorf("invalid format for header - size %d", len(header))
		}
		request.Header.Set(header[0], header[1])
	}
	return nil
}

func (c *HttpClient) addQueryParams(request *http.Request, queryParams string) error {
	queryParams = strings.TrimSpace(queryParams)
	if queryParams == "" {
		return nil
	}
	urlQuery := request.URL.Query()
	for _, line := range strings.Split(queryParams, "\n") {
		queryParam := strings.Split(line, "=")
		if len(queryParam) != 2 {
			return fmt.Errorf("invalid format for query params - size %d", len(queryParam))
		}
		urlQuery.Set(queryParam[0], queryParam[1])
	}
	return nil
}
