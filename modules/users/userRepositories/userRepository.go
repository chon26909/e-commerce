package userRepositories

import (
	"github.com/chon26909/e-commerce/modules/users"
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
	panic("unimplemented")
}
