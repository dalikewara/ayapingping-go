package repository

type User interface {
	// FindAll finds all user data from database.
	FindAll(param UserFindAllParam) UserFindAllResult
}

type Role interface {
	// FindByUserID finds role data by user id from database.
	FindByUserID(param RoleFindByUserIDParam) RoleFindByUserIDResult
}

type Repository struct {
	User User
	Role Role
}

// New generates new repository.
func New(param NewParam) *Repository {
	userRepo := NewUserMySQL(NewUserMySQLParam{
		DB: param.MySQLDB,
	})

	roleRepo := NewRoleMySQL(NewRoleMySQLParam{
		DB: param.MySQLDB,
	})

	return &Repository{
		User: userRepo,
		Role: roleRepo,
	}
}
