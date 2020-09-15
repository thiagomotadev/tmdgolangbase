package tmdgolangbase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router ...
type Router struct {
	MuxRouter *mux.Router
	path      string
}

// ListenAndServe ...
func (router *Router) ListenAndServe() {
	log.Fatal(http.ListenAndServe(":80", router.MuxRouter).Error())
}

func (router *Router) setInitialRouterPrefix() {
	if router.path == "" {
		router.path = "/api"
	}
}

// Subrouter ...
func (router *Router) Subrouter(path string) (subrouter *Router) {
	router.setInitialRouterPrefix()
	subrouter = &Router{
		MuxRouter: router.MuxRouter,
		path:      fmt.Sprintf("%v%v", router.path, path),
	}
	return
}

// AddRoute ...
func (router *Router) AddRoute(path string, method string, fHandler func(RequestAndResponse) error) {
	router.setInitialRouterPrefix()
	router.MuxRouter.Path(fmt.Sprintf("%v%v", router.path, path)).Methods(method).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rar := RequestAndResponse{
			w: w,
			r: r,
		}
		err := fHandler(rar)
		if err != nil {
			rar.setResponseError(err)
		}
	})
}

// AddRouteOnRoot ...
func (router *Router) AddRouteOnRoot(method string, fHandler func(RequestAndResponse) error) {
	router.AddRoute("", method, fHandler)
}

// AddRouteOnID ...
func (router *Router) AddRouteOnID(method string, fHandler func(RequestAndResponse, uint) error) {
	router.AddRoute("/{id:[0-9]+}", method, func(rar RequestAndResponse) (err error) {
		id, err := rar.GetRequestUintParameter("id")
		if err != nil {
			return
		}
		err = fHandler(rar, id)
		return
	})
}

// AddRouteOnUUID ...
func (router *Router) AddRouteOnUUID(method string, fHandler func(RequestAndResponse, string) error) {
	router.AddRoute("/{uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", method, func(rar RequestAndResponse) (err error) {
		uuid, err := rar.GetRequestParameter("uuid")
		if err != nil {
			return
		}
		err = fHandler(rar, uuid)
		return
	})
}
