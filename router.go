package bone

import (
	"log"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{Router: mux.NewRouter()}
}

func (r *Router) RegisterEndpoint(method, path string, e endpoint.Endpoint, dec http.DecodeRequestFunc) *Router {
	method = strings.ToUpper(method)
	log.Printf("Registering Endpoint: %s %s\n", method, path)
	r.Methods(method).Path(path).Handler(http.NewServer(e, dec, Http.EncodeJSON))
	return r
}
