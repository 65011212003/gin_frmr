package repository

import (
	"gin_frmr/internal/domain"
	"time"

	"gorm.io/gorm"
)

// UserModel is the GORM model for users table
type UserModel struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	var models []UserModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, len(models))
	for i, m := range models {
		users[i] = toDomainUser(m)
	}
	return users, nil
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var model UserModel
	if err := r.db.First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	user := toDomainUser(model)
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	model := toUserModel(user)
	if err := r.db.Create(&model).Error; err != nil {
		return err
	}
	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *userRepository) Update(user *domain.User) error {
	model := toUserModel(user)
	if err := r.db.Save(&model).Error; err != nil {
		return err
	}
	user.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&UserModel{}, id).Error
}

// Helper functions for mapping
func toDomainUser(m UserModel) domain.User {
	return domain.User{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func toUserModel(u *domain.User) UserModel {
	return UserModel{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
