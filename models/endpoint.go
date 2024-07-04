package models

import "fmt"

var ENDPOINT_METHODS []string = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"PATCH",
	"HEAD",
	"CONNECT",
	"OPTIONS",
	"TRACE",
}

var ENDPOINT_METHOD_COLOR map[string]string = map[string]string{
	"GET":     "blue",
	"POST":    "green",
	"PUT":     "orange",
	"DELETE":  "red",
	"PATCH":   "darkcyan",
	"HEAD":    "darkcyan",
	"CONNECT": "darkcyan",
	"OPTIONS": "darkcyan",
	"TRACE":   "darkcyan",
}

type Endpoint struct {
	Id     int
	Method string
	Url    string

	RequestQueryParams string
	RequestHeaders     string
	RequestBodyType    string
	RequestBody        string

	ResponseBodyType string
	ResponseBody     string
	StatusCode       int
}

func (e *Endpoint) MethodIndex() int {
	for i, method := range ENDPOINT_METHODS {
		if e.Method == method {
			return i
		}
	}
	return -1
}

func (e *Endpoint) String() string {
	return fmt.Sprintf("[%s]%s[white] %s", ENDPOINT_METHOD_COLOR[e.Method], e.Method, e.Url)
}
