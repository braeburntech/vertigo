package hello

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, World!"))
}

func Init(r *mux.Route) {
	s := r.Subrouter()
	s.HandleFunc("/", Hello).Name("hello")
}
