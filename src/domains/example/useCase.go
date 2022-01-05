package example

type UseCase struct {
	exampleService ServiceInterface
	// If you need to use other services, you can pass them here.
	// Example:
	// EmailService EmailServiceInterface
	// SMSService   SMSServiceInterface
}

// NewUseCase generates new UseCase.
func NewUseCase(param NewUseCaseParam) UseCaseInterface {
	return &UseCase{
		exampleService: param.ExampleService,
		// If you need to use other services, you can pass them here.
		// Example:
		// emailService: param.EmailServiceInterface,
		// sMSService:   param.SMSServiceInterface,
	}
}

// GetAndChangeName gets example data and changes its name.
func (u *UseCase) GetAndChangeName(param UseCaseGetAndChangeNameParam) UseCaseGetAndChangeNameResult {
	resGet := u.exampleService.Get(ServiceGetParam{
		Id: param.Id,
	})
	ex := resGet.Example
	isChange := false
	resUpdateName := u.exampleService.UpdateName(ServiceUpdateNameParam{
		Id:   ex.Id,
		Name: param.Name,
	})
	if resUpdateName.Error == nil {
		isChange = true
		ex.Name = param.Name
	}
	return UseCaseGetAndChangeNameResult{
		Example:   ex,
		IsChanged: isChange,
		Error:     nil,
	}
}
