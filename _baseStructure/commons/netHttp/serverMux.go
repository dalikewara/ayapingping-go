package netHttp

import "net/http"

func ServerMux() *http.ServeMux {
	return http.NewServeMux()
}
