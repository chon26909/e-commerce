package userRepositories

import (
	"fmt"

	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userPatterns"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	InsertUser(req *users.UserRegisterRequest, isAdmin bool) (*users.UserPassport, error)
	FindOneUserbyEmail(email string) (*users.UserCredentialCheck, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepositories(db *sqlx.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// InsertUser implements IUserRepository.
func (r *userRepository) InsertUser(req *users.UserRegisterRequest, isAdmin bool) (*users.UserPassport, error) {

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

	user, err := result.Result()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindOneUserbyEmail(email string) (*users.UserCredentialCheck, error) {

	query := `
		SELECT 
			"id",
			"email",
			"password",
			"username",
			"role_id",
		FROM "users"
		WHERE "email" = $1;
	`

	user := new(users.UserCredentialCheck)
	if err := r.db.Get(user, query, email); err != nil {
		return nil, fmt.Errorf("user not found : %v", err)
	}

	return user, nil
}
