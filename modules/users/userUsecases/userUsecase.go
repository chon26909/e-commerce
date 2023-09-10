package userUsecases

import (
	"fmt"

	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userRepositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	InsertCustomer(req *users.UserRegisterRequest) (*users.UserPassport, error)
	GetPassport(req *users.UserCredential) (*users.UserPassport, error)
}

type userUsecase struct {
	config         *config.IConfig
	userRepository userRepositories.IUserRepository
}

func NewUserUsecase(config *config.IConfig, userRepository userRepositories.IUserRepository) IUserUsecase {
	return &userUsecase{config: config, userRepository: userRepository}
}

func (u *userUsecase) InsertCustomer(req *users.UserRegisterRequest) (*users.UserPassport, error) {

	// hashing a password
	if err := req.BcryptHash(); err != nil {
		return nil, err
	}

	// insert user
	result, err := u.userRepository.InsertUser(req, false)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userUsecase) GetPassport(req *users.UserCredential) (*users.UserPassport, error) {
	user, err := u.userRepository.FindOneUserbyEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("failed to compare password")
	}

	passport := &users.UserPassport{
		User: &users.User{
			Id:       user.Id,
			Email:    user.Email,
			UserName: user.Username,
			RoleId:   user.RoleId,
		},
		Token: nil,
	}

	return passport, nil
}
