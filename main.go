package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/PhilShaughnes/snipboxgo/internal/web"
)

var (
	port = flag.Int("port", 4321, "server port")
)

func main() {
	flag.Parse()

	r := web.NewRouter()

	middleware := web.Use(web.Logging)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: middleware(r),
	}

	slog.Info("starting server", slog.Int("port", *port))
	server.ListenAndServe()
}
