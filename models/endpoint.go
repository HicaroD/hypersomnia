package models

import "fmt"

var METHOD_COLOR map[string]string = map[string]string {
	"GET": "blue",
	"POST": "green",
	"PUT": "orange",
	"DELETE": "red",
	"PATCH": "darkcyan",
	"HEAD": "darkcyan",
	"CONNECT": "darkcyan",
	"OPTIONS": "darkcyan",
	"TRACE": "darkcyan",
}

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

func (e *Endpoint) String() string {
	return fmt.Sprintf("[%s]%s[white] %s", METHOD_COLOR[e.Method], e.Method, e.Url)
}
