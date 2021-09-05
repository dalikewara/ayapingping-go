package example

// APITestRequest request.
type APITestRequest struct {}

// APITestResponse response.
type APITestResponse struct {
	Code string
	Message string
	Data *APITestResponseData
}

// APITestResponseData response data.
type APITestResponseData struct {
	Id int64
	Name string
	IsChanged bool
}
