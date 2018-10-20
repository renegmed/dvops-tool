package nap

import (
	"net/http"
	"strings"
)

// Client doesn't keep authentication information. It is delegated to
// Authentication entity.
type Client struct {
	Client   *http.Client
	AuthInfo Authentication
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) SetAuth(auth Authentication) {
	c.AuthInfo = auth
}

// ProcessRequest send a request to a resource (web api service)
func (c *Client) ProcessRequest(baseURL string, res *RestResource, params map[string]string) error {

	// form URL

	// TrimLeft returns a slice of the string with all leading
	// Unicode code points contained in '/' removed.
	endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")

	// TrimRight returns a slice of the string, with all trailing
	// Unicode code points contained in "/" removed.
	trimmedBaseURL := strings.TrimRight(baseURL, "/")
	url := trimmedBaseURL + "/" + endpoint

	// form a request
	req, err := http.NewRequest(res.Method, url, nil) // nil - no payload it passed
	if err != nil {
		return err
	}

	// if authorization is required, add authorization information to request header
	if c.AuthInfo != nil {
		req.Header.Add("Authorization", c.AuthInfo.AuthorizationHeader())
	}

	// Call the API service and get the response
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if err := res.Router.CallFunc(resp); err != nil {
		return err
	}

	return nil
}
