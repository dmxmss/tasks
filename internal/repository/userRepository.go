package repository

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"gorm.io/gorm"

	"errors"
)

type UserRepository interface {
	CreateUser(entities.CreateUserDto) (*entities.User, error)
	GetUserByEmail(string) (*entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(createUser entities.CreateUserDto) (*entities.User, error) {
	user := entities.User{
		FullName: createUser.FullName,
		Email: createUser.Email,
		Password: createUser.Password, // TODO hashing
		City: createUser.City,
	}

	if err := ur.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, e.ErrUserAlreadyExists
		} else {
			return nil, e.ErrDbTransactionFailed
		}
	}

	return &user, nil	
}

func (ur *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserNotFound
		} else {
			return nil, e.ErrDbTransactionFailed
		}
	}

	return &user, nil
}
