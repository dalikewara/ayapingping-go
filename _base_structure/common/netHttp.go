package common

import "net/http"

func NewNetHttpServer() *http.Server {
	return &http.Server{}
}

func NewNetHttpServerMux() *http.ServeMux {
	return http.NewServeMux()
}
