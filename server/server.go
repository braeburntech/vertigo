package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/braeburntech/vertigo/hello"
)

func main() {
	r := mux.NewRouter()

	hello.Init(r.PathPrefix("/hello"))

	http.Handle("/", r)

	log.Print(http.ListenAndServe(":8080", nil))
}
