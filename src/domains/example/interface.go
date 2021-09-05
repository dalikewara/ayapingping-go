package example

// RepositoryInterface interface.
type RepositoryInterface interface {
	// FindById finds `example` from database by id.
	FindById(param RepositoryFindByIdParam) RepositoryFindByIdResult

	// UpdateNameById updates `example` name to database by id.
	UpdateNameById(param RepositoryUpdateNameByIdParam) RepositoryUpdateNameByIdResult
}

// ServiceInterface interface.
type ServiceInterface interface {
	// Get gets `example` data.
	Get(param ServiceGetParam) ServiceGetResult

	// UpdateName updates `example` data.
	UpdateName(param ServiceUpdateNameParam) ServiceUpdateNameResult
}

// UseCaseInterface interface.
type UseCaseInterface interface {
	// GetAndChangeName gets `example` data then changes its name.
	GetAndChangeName(param UseCaseGetAndChangeNameParam) UseCaseGetAndChangeNameResult
}
