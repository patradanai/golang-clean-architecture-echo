package userrepo

type (
	UserRepository struct {
	}

	IUserRepository interface {
		FindOne()
	}
)

func (s *UserRepository) FindOne() {
}
