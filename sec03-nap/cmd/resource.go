package nap

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

type RestResource struct {
	Endpoint string    // e.g. /get, /get/{.user} thus this resource can be templated
	Method   string    // e.g. GET
	Router   *CBRouter // different router for each resource. Pointer
	// is used so other resources can share the same router
}

func NewResource(endpoint, method string, router *CBRouter) *RestResource {
	return &RestResource{
		Endpoint: endpoint,
		Method:   method,
		Router:   router,
	}
}

func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}

	// NOTE: Template is a specialized Template from "text/template" that produces a safe
	// HTML document fragment.

	// make endpoint into template
	t, err := template.New("resource").Parse(r.Endpoint)
	if err != nil {
		log.Fatalln("Unable to parse endpoint")
	}
	// Create empty bytes buffer struct
	buffer := &bytes.Buffer{}

	// Fill in the resource templates with values from params
	t.Execute(buffer, params)

	// read bytes buffer and returns array of bytes
	endpoint, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalln("Unable to read endpoint")
	}

	// convert bytes buffer into string
	return string(endpoint)

}
