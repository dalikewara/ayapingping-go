package ping

import (
	"net/http"
)

type V1HttpServer struct{}

// String handles string response
func (v *V1HttpServer) String(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
	return
}

// New

func NewV1HttpServer() *V1HttpServer {
	return &V1HttpServer{}
}
