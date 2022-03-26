package ginGonic

// Routes handles http routes.
func (h *Handler) Routes() {
	h.router.GET("/api/v1/user/get-all", h.APIGetAll)
	h.router.POST("/api/v1/user/create", h.APICreate)
}
