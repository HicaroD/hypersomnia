package main

type Endpoint struct {
	Id     int
	Method string
	Url    string

	QueryParams     string
	Headers         string
	RequestBodyType string
	RequestBody     string

	ResponseBodyType string
	ResponseBody     string
	StatusCode       int
}
