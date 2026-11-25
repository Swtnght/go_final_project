package server

import (
	"net/http"
	"strconv"

	"github.com/Swtnght/go_final_project/pkg/api"
)

const (
	Port   = 7540
	webDir = "web"
)

func Run() error {
	mux := http.NewServeMux()
	mux.Handle("/api/", api.Init())
	mux.Handle("/", http.FileServer(http.Dir(webDir)))
	return http.ListenAndServe(":"+strconv.Itoa(Port), mux)
}
