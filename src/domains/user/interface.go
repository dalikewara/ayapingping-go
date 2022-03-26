package user

// UseCaseInterface is the use case interface contract of the domain.
// UseCase describes what the domain can do based on the main business flow.
// UseCase is a place where you use & call service functions of the domain
// or services from other domains
type UseCaseInterface interface {
	GetAll(param UseCaseGetAllParam) UseCaseGetAllResult
	Create(param UseCaseCreateParam) UseCaseCreateResult
}

// ServiceInterface is the service interface contract of the domain.
// Service is the smallest partial from what the domain can do based on the main business flow.
// Service is a place where you use & call repository functions of the domain.
type ServiceInterface interface {
	GetAll(param ServiceGetAllParam) ServiceGetAllResult
	GetByUsername(param ServiceGetByUsernameParam) ServiceGetByUsernameResult
	Create(param ServiceCreateParam) ServiceCreateResult
}

// RepositoryInterface is the repository interface contract of the domain.
// Repository is a place where you communicate with the real external data source or database.
type RepositoryInterface interface {
	FindAll(param RepositoryFindAllParam) RepositoryFindAllResult
	FindByUsername(param RepositoryFindByUsernameParam) RepositoryFindByUsernameResult
	Insert(param RepositoryInsertParam) RepositoryInsertResult
}
