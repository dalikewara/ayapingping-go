package example

// HttpResponse is the response model of the all incoming http requests in delivery/http.
type HttpResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HttpTestRequest is the payload model of the test request in delivery/http.
type HttpTestRequest struct{}

// HttpTestResponseData is the response data model of the test request in delivery/http.
type HttpTestResponseData struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	IsChanged bool   `json:"is_changed"`
}
