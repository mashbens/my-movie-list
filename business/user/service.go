package user

import (
	"errors"
	"log"

	"github.com/mashbens/my-movie-list/business/user/dto"
	"github.com/mashbens/my-movie-list/business/user/entity"
	_user "github.com/mashbens/my-movie-list/business/user/response"

	// repo "github.com/mashbens/my-movie-list/repository/user"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByUserID(userID string) (entity.User, error)
}

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) (*_user.UserResponse, error)
	UpdateUser(updateUserRequest dto.UpdateUserRequest) (*_user.UserResponse, error)
	FindUserByEmail(email string) (*_user.UserResponse, error)
	FindUserByID(userID string) (*_user.UserResponse, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) UpdateUser(updateUserRequest dto.UpdateUserRequest) (*_user.UserResponse, error) {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&updateUserRequest))

	if err != nil {
		return nil, err
	}

	user, err = c.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	res := _user.NewUserResponse(user)
	return &res, nil

}

func (c *userService) CreateUser(registerRequest dto.RegisterRequest) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByEmail(registerRequest.Email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	user, _ = c.userRepo.InsertUser(user)
	res := _user.NewUserResponse(user)
	return &res, nil

}

func (c *userService) FindUserByEmail(email string) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	userResponse := _user.NewUserResponse(user)
	return &userResponse, nil
}

func (c *userService) FindUserByID(userID string) (*_user.UserResponse, error) {
	user, err := c.userRepo.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	userResponse := _user.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}
