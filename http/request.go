package http

type Request struct {
	Method, Url, Body, QueryParams, Headers string
}
