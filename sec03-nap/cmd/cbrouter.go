package nap

import (
	"fmt"
	"net/http"
)

// RouterFunc has content interface{} parameter that can be used
// to pass a value or content back to the caller i.e. callback function
type RouterFunc func(resp *http.Response, content interface{}) error

// CBRouter maintains a list of routers. If router has no router function
// a default router will be called
type CBRouter struct {
	Routers       map[int]RouterFunc // int are the status code e.g. 200, 404, etc
	DefaultRouter RouterFunc         // if no status code is set, this router will be called
}

func NewRouter() *CBRouter {
	return &CBRouter{
		Routers: make(map[int]RouterFunc),
		DefaultRouter: func(resp *http.Response, _ interface{}) error {
			return fmt.Errorf("From: %s received unknown status: %d",
				resp.Request.URL.String(), resp.StatusCode)
		},
	}
}

// RegisterFunc will register a function with a status
func (r *CBRouter) RegisterFunc(status int, fn RouterFunc) {
	r.Routers[status] = fn
}

// CallFunc calls a registered function in the router
// content parameter is optional
func (r *CBRouter) CallFunc(resp *http.Response, content ...interface{}) error { // parameter optional
	fn, ok := r.Routers[resp.StatusCode]
	if !ok {
		fn = r.DefaultRouter
	}
	if err := fn(resp, content); err != nil {
		return err
	}
	return nil
}
