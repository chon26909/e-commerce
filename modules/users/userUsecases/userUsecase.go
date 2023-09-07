package userUsecases

import (
	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userRepositories"
)

type IUserUsecase interface {
	InsertCustomer(req *users.UserRegisterRequest) (*users.UserPassport, error)
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
