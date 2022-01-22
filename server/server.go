package server

import (
	"net/http"
)

//goland:noinspection GoUnusedExportedFunction
func NewStaticServerMux(path string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
	return mux
}
