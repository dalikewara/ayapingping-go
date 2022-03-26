package user

import "context"

type NewServiceParam struct {
	Repository RepositoryInterface
}

type ServiceGetAllParam struct {
	Ctx context.Context // If you want to use context.
}

type ServiceGetAllResult struct {
	Users *[]User `json:"users"`
	Error error   `json:"error"`
}

type ServiceGetByUsernameParam struct {
	Username string          `json:"username"`
	Ctx      context.Context // If you want to use context.
}

type ServiceGetByUsernameResult struct {
	User  *User `json:"user"`
	Error error `json:"error"`
}

type ServiceCreateParam struct {
	Username string          `json:"username"`
	Ctx      context.Context // If you want to use context.
}

type ServiceCreateResult struct {
	Error error `json:"error"`
}
