package bone

import (
	"github.com/gorilla/mux"
)

// Router is a wrapper of *mux.Router.
type Router struct {
	*mux.Router
}

// NewRouter returns a Router instance of Bone Application.
func NewRouter() *Router {
	return &Router{Router: mux.NewRouter()}
}
