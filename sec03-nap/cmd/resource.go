package nap

type RestResource struct {
	Endpoint string    // e.g. /get, /get/{.user}
	Method   string    // e.g. GET
	Router   *CBRouter // different router for each resource. pointer is used so other resources can share the same router
}
