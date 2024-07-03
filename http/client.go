package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HttpClient struct {
	client *http.Client
}

func New(client *http.Client) *HttpClient {
	return &HttpClient{client}
}

func (c *HttpClient) DoRequest(request Request) (*Response, error) {
	req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))
	if err != nil {
		return nil, err
	}

	err = c.addQueryParams(req, request.QueryParams)
	if err != nil {
		return nil, err
	}

	err = c.addHeaders(req, request.Headers)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	responseAsStr, err := c.responseToString(response)
	if err != nil {
		return nil, err
	}

	return &Response{responseAsStr}, nil
}

func (c *HttpClient) addHeaders(request *http.Request, headers string) error {
	headers = strings.TrimSpace(headers)
	if headers == "" {
		return nil
	}

	newHeaders := request.Header
	for _, line := range strings.Split(headers, "\n") {
		header := strings.Split(line, "=")
		if len(header) != 2 {
			return fmt.Errorf("invalid format for header - size %d", len(header))
		}
		newHeaders.Add(header[0], header[1])
	}
	request.Header = newHeaders
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
		urlQuery.Add(queryParam[0], queryParam[1])
	}
	request.URL.RawQuery = urlQuery.Encode()
	return nil
}

// TODO: deal with more types of response, not only JSON
func (c *HttpClient) responseToString(response *http.Response) (string, error) {
	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	formattedJsonBuffer := &bytes.Buffer{}
	err = json.Indent(formattedJsonBuffer, respBytes, "", "  ")
	if err != nil {
		return "", err
	}

	return formattedJsonBuffer.String(), nil
}
