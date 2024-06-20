package web

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", Root)
	// router.HandleFunc("GET /snippet/{id}", show)
	router.HandleFunc("GET /snippet", Show)
	router.HandleFunc("GET /snippet/create", Create)

	fileServer := http.FileServer(http.Dir("./internal/ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	r := http.NewServeMux()
	r.HandleFunc("GET /hello/{name}", hello)
	router.Handle("/v1/", http.StripPrefix("/v1", r))

	return router
}

func hello(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("name")
	w.Write([]byte("hello " + v))
}
