// Go Api server
// @jeffotoni

package cf

type EndpointS struct {
	Ping     string
	Sendjson string
	Produto  string
	PostAuth string
}

func Endpoint() *EndpointS {
	return &EndpointS{
		Ping:     "/ping",     // POST/GET
		Sendjson: "/sendjson", // POST
		Produto:  "/produto",  // [GET/PUT/POST/DELETE]
		PostAuth: "/auth",     // [POST]
	}
}
