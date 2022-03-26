package user

type Service struct {
	repository RepositoryInterface
}

// NewService generates new Service.
func NewService(param NewServiceParam) ServiceInterface {
	return &Service{
		repository: param.Repository,
	}
}

// GetAll gets all user data.
func (s *Service) GetAll(param ServiceGetAllParam) ServiceGetAllResult {
	result := ServiceGetAllResult{}
	users := s.repository.FindAll(RepositoryFindAllParam{
		Ctx: param.Ctx,
	})
	if users.Error != nil {
		result.Error = users.Error
		return result
	}
	result.Users = users.Users
	return result
}

// GetByUsername gets user data by username.
func (s *Service) GetByUsername(param ServiceGetByUsernameParam) ServiceGetByUsernameResult {
	result := ServiceGetByUsernameResult{}
	user := s.repository.FindByUsername(RepositoryFindByUsernameParam{
		Username: param.Username,
		Ctx:      param.Ctx,
	})
	if user.Error != nil {
		result.Error = user.Error
		return result
	}
	result.User = user.User
	return result
}

// Create creates user data.
func (s *Service) Create(param ServiceCreateParam) ServiceCreateResult {
	result := ServiceCreateResult{}
	reply := s.repository.Insert(RepositoryInsertParam{
		Username: param.Username,
		Ctx:      param.Ctx,
	})
	if reply.Error != nil {
		result.Error = reply.Error
		return result
	}
	return result
}
