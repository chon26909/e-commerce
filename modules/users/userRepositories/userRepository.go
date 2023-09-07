package userRepositories

import (
	"fmt"

	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userPatterns"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	InsertUser(req *users.UserRegisterRequest, isAdmin bool) (*users.UserPassport, error)
}

type userRepositories struct {
	db *sqlx.DB
}

func NewUserRepositories(db *sqlx.DB) IUserRepository {
	return &userRepositories{
		db: db,
	}
}

// InsertUser implements IUserRepository.
func (r *userRepositories) InsertUser(req *users.UserRegisterRequest, isAdmin bool) (*users.UserPassport, error) {

	result := userPatterns.InsertUser(r.db, req, isAdmin)

	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("result: ", result)

	user, err := result.Result()
	if err != nil {
		return nil, err
	}

	return user, nil
}
