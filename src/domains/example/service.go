package example

type Service struct {
	repository RepositoryInterface
}

// NewService generates new Service.
func NewService(param NewServiceParam) ServiceInterface {
	return &Service{
		repository: param.Repository,
	}
}

// Get gets example data.
func (s *Service) Get(param ServiceGetParam) ServiceGetResult {
	res := s.repository.FindById(RepositoryFindByIdParam{
		Id: param.Id,
	})
	return ServiceGetResult{
		Example: res.Example,
		Error:   nil,
	}
}

// UpdateName updates example name.
func (s *Service) UpdateName(param ServiceUpdateNameParam) ServiceUpdateNameResult {
	_ = s.repository.UpdateNameById(RepositoryUpdateNameByIdParam{
		Id:   param.Id,
		Name: param.Name,
	})
	return ServiceUpdateNameResult{
		Error: nil,
	}
}
