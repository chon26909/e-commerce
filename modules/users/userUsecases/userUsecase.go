package userUsecases

import (
	"fmt"

	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userRepositories"
	"github.com/chon26909/e-commerce/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	InsertCustomer(req *users.UserRegisterRequest) (*users.UserPassport, error)
	GetPassport(req *users.UserCredential) (*users.UserPassport, error)
	RefreshPassport(req *users.UserRefreshCredential) (*users.UserPassport, error)
}

type userUsecase struct {
	config         config.IConfig
	userRepository userRepositories.IUserRepository
}

func NewUserUsecase(config config.IConfig, userRepository userRepositories.IUserRepository) IUserUsecase {
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

	accessToken, err := auth.NewAuth(string(auth.Access), u.config.Jwt(), &users.UserClaims{
		Id:     user.Id,
		RoleId: user.RoleId,
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.NewAuth(string(auth.Refresh), u.config.Jwt(), &users.UserClaims{
		Id:     user.Id,
		RoleId: user.RoleId,
	})
	if err != nil {
		return nil, err
	}

	passport := &users.UserPassport{
		User: &users.User{
			Id:       user.Id,
			Email:    user.Email,
			UserName: user.Username,
			RoleId:   user.RoleId,
		},
		Token: &users.UserToken{
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken.SignToken(),
		},
	}

	if err := u.userRepository.InsertOAuth(passport); err != nil {
		return nil, err
	}

	return passport, nil
}

func (u *userUsecase) RefreshPassport(req *users.UserRefreshCredential) (*users.UserPassport, error) {

	claims, err := auth.ParseToken(u.config.Jwt(), req.RefreshToken)
	if err != nil {
		return nil, err
	}

	oauth, err := u.userRepository.FindOneOauth(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	profile, err := u.userRepository.GetProfile(oauth.UserId)
	if err != nil {
		return nil, err
	}

	newClaims := &users.UserClaims{
		Id:     profile.Id,
		RoleId: profile.RoleId,
	}

	accessToken, err := auth.NewAuth(string(auth.Access), u.config.Jwt(), newClaims)
	if err != nil {
		return nil, err
	}

	refreshToken := auth.RepeatToken(u.config.Jwt(), newClaims, claims.ExpiresAt.Unix())
	if err != nil {
		return nil, err
	}

	passport := &users.UserPassport{
		User: profile,
		Token: &users.UserToken{
			Id:           oauth.Id,
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken,
		},
	}

	if err := u.userRepository.UpdateOauth(passport.Token); err != nil {
		return nil, err
	}

	return passport, nil
}
