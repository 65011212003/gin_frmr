package usecase

import "gin_frmr/internal/domain"

// UserUseCase defines the interface for user business logic
type UserUseCase interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	CreateUser(name, email string) (*domain.User, error)
	UpdateUser(id uint, name, email string) (*domain.User, error)
	DeleteUser(id uint) error
}

type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(repo domain.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) GetAllUsers() ([]domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) GetUserByID(id uint) (*domain.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

func (u *userUseCase) CreateUser(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, domain.ErrInvalidInput
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}

	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) UpdateUser(id uint, name, email string) (*domain.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	if err := u.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) DeleteUser(id uint) error {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrUserNotFound
	}

	return u.userRepo.Delete(id)
}
