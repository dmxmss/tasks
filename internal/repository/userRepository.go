package repository

import (
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"
	"gorm.io/gorm"

	"errors"
)

type UserRepository interface {
	CreateUser(entities.CreateUserDto) (*entities.User, error)
	GetUserBy(entities.SearchUserDto) (*entities.User, error)
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
		Password: createUser.Password,
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

func (ur *userRepository) GetUserBy(searchUser entities.SearchUserDto) (*entities.User, error) {
	var user entities.User
	query := ur.db.Model(&entities.User{})

	if searchUser.ID != nil {
		query = query.Where("id = ?", *searchUser.ID)		
	}

	if searchUser.Email != nil {
		query = query.Where("email = ?", *searchUser.Email)		
	}

	query.First(&user)

	return &user, nil
}
