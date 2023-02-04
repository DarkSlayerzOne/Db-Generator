package router

import "net/http"

type MuxRouter interface {
	SERVE(port string)
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	PATCH(uri string, f func(w http.ResponseWriter, r *http.Request))
}
