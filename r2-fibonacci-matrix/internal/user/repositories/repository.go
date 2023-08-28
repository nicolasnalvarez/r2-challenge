package repositories

import (
	"gorm.io/gorm"
	"r2-fibonacci-matrix/internal/user/entities"
)

type (
	Repository struct {
		db *gorm.DB
	}

	UserRepository interface {
		SaveUser(user entities.User) error
		FindUserByEmail(email string) (entities.User, error)
	}
)

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) SaveUser(user entities.User) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r Repository) FindUserByEmail(email string) (entities.User, error) {
	var user entities.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}
