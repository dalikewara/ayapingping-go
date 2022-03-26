package user

import "errors"

type UseCase struct {
	service ServiceInterface
	// If you need to use other services from other domains, you can pass them here.
	// Example:
	// EmailService EmailServiceInterface
	// SMSService   SMSServiceInterface
}

// NewUseCase generates new UseCase.
func NewUseCase(param NewUseCaseParam) UseCaseInterface {
	return &UseCase{
		service: param.Service,
		// If you need to use other services from other domains, you can pass them here.
		// Example:
		// emailService: param.EmailServiceInterface,
		// sMSService:   param.SMSServiceInterface,
	}
}

// GetAll gets all user data.
func (u *UseCase) GetAll(param UseCaseGetAllParam) UseCaseGetAllResult {
	result := UseCaseGetAllResult{}
	users := u.service.GetAll(ServiceGetAllParam{
		Ctx: param.Ctx,
	})
	if users.Error != nil {
		result.Error = users.Error
		return result
	}
	result.Users = users.Users
	return result
}

// Create creates user data.
func (u *UseCase) Create(param UseCaseCreateParam) UseCaseCreateResult {
	result := UseCaseCreateResult{}

	// Check user exists.
	exists := u.service.GetByUsername(ServiceGetByUsernameParam{
		Username: param.Username,
		Ctx:      param.Ctx,
	})
	if exists.Error != nil {
		result.Error = exists.Error
		return result
	}
	if exists.User != nil {
		result.Error = errors.New("user already exists")
		return result
	}

	reply := u.service.Create(ServiceCreateParam{
		Username: param.Username,
		Ctx:      param.Ctx,
	})
	if reply.Error != nil {
		result.Error = reply.Error
		return result
	}
	return result
}
