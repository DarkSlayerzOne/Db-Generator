package router

import (
	"Db-Generator/src/core/ports/router"
	"Db-Generator/src/pkg"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func MuxRouterImpl() router.MuxRouter {
	return &muxRouter{}
}

func (*muxRouter) SERVE(port string) {
	log.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("invoke GET request ")
	muxDispatcher.HandleFunc(uri, f).Methods(pkg.GET)
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("invoke POST request")
	muxDispatcher.HandleFunc(uri, f).Methods(pkg.POST)
}

func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("invoke PUT request")
	muxDispatcher.HandleFunc(uri, f).Methods(pkg.PUT)
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("invoke DELETE request")
	muxDispatcher.HandleFunc(uri, f).Methods(pkg.DELETE)
}

func (*muxRouter) PATCH(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	log.Println("invoke PATCH request")
	muxDispatcher.HandleFunc(uri, f).Methods(pkg.PATCH)
}
