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
	api.Init()
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	return http.ListenAndServe(":"+strconv.Itoa(Port), nil)
}
