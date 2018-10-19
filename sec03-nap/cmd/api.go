package nap

type API struct {
	BaseURL       string                  // e.g. https://httpbin.org
	Resources     map[string]RestResource // an api registry thus becomes a multiple api caller
	DefaultRouter *CBRouter               // callback default router to be used if non-registered router is called
	Client        *Client                 // nap client, an api caller
}
