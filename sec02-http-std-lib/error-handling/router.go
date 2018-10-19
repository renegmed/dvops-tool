package main

import (
	"log"
	"net/http"
)

type HandlerFunc func(*http.Response)

type ResponseRouter struct {
	Handlers       map[int]HandlerFunc
	DefaultHandler HandlerFunc
}

func ErrorHandlerFunc(r *http.Response) {
	log.Fatalln("\t[ErrorHandlerFunc] Unhandled Response: ", r.StatusCode)
}

func NewRouter() *ResponseRouter {
	return &ResponseRouter{
		Handlers:       make(map[int]HandlerFunc),
		DefaultHandler: ErrorHandlerFunc,
	}
}

func (r *ResponseRouter) Register(status int, handler HandlerFunc) {
	r.Handlers[status] = handler
}

func (r *ResponseRouter) Process(resp *http.Response) {
	f, ok := r.Handlers[resp.StatusCode]
	if !ok {
		r.DefaultHandler(resp)
		return
	}
	f(resp)
}
