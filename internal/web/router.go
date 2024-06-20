package web

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", root)
	router.HandleFunc("GET /snippet/{id}", show)
	router.HandleFunc("GET /snippet", show)
	router.HandleFunc("GET /snippet/create", create)

	r := http.NewServeMux()
	r.HandleFunc("GET /hello/{name}", hello)
	router.Handle("/v1/", http.StripPrefix("/v1", r))

	return router
}

func hello(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("name")
	w.Write([]byte("hello " + v))
}

func show(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("id")
	w.Write([]byte("snippet: " + v))
}

func create(w http.ResponseWriter, _ *http.Request) {
	// v := r.PathValue("name")
	w.Write([]byte("creating snippet"))
}

func root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("this is snipbox"))
}
