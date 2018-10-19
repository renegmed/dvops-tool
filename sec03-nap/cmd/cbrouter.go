package nap

import "net/http"

type RouterFunc func(client *http.Client, content interface{})

type CBRouter struct {
	Routers       map[int]RouterFunc // int are the status code e.g. 200, 404, etc
	DefaultRouter RouterFunc         // if no status code is set, this router will be called
}
