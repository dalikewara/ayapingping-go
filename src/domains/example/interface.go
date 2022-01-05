package example

// RepositoryInterface is the repository interface contract of the domain.
// Repository is a place where you communicate with the real database.
type RepositoryInterface interface {
	// FindById finds example data from database by id.
	FindById(param RepositoryFindByIdParam) RepositoryFindByIdResult

	// UpdateNameById updates example name on database by id.
	UpdateNameById(param RepositoryUpdateNameByIdParam) RepositoryUpdateNameByIdResult
}

// ServiceInterface is the service interface contract of the domain.
// Service is a place where you use & call repository functions of the domain,
// or libraries outside the domain.
type ServiceInterface interface {
	// Get gets example data.
	Get(param ServiceGetParam) ServiceGetResult

	// UpdateName updates example name.
	UpdateName(param ServiceUpdateNameParam) ServiceUpdateNameResult
}

// UseCaseInterface is the use case interface contract of the domain.
// UseCase is a place where you use & call service functions of the domain,
// or services from other domains.
type UseCaseInterface interface {
	// GetAndChangeName gets example data and changes its name.
	GetAndChangeName(param UseCaseGetAndChangeNameParam) UseCaseGetAndChangeNameResult
}
