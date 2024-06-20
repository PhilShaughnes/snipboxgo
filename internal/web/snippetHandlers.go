package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"
)

func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles(
		"./internal/ui/html/home.page.html",
		"./internal/ui/html/base.layout.html",
		"./internal/ui/html/footer.partial.html",
	)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "snippet: %d", id)
}

func Create(w http.ResponseWriter, _ *http.Request) {
	// v := r.PathValue("name")
	w.Write([]byte("creating snippet"))
}
