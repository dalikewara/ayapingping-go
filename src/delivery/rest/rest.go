package rest

type HandlerFunc func(...interface{}) interface{}

type Server interface {
	// RegisterRoutes registers REST routes.
	RegisterRoutes()
	// Serve serves REST application.
	Serve()

	// APIUserV1GetAll handles REST API to get all user data.
	APIUserV1GetAll() HandlerFunc

	// APIRoleV1GetByUserID handles REST API to get role data by user id.
	APIRoleV1GetByUserID() HandlerFunc
}
