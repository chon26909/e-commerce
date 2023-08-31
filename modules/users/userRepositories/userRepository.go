package userRepositories

import "github.com/jmoiron/sqlx"

type IUserRepository interface {
}

type userRepositories struct {
	db *sqlx.DB
}

func NewUserRepositories(db *sqlx.DB) IUserRepository {
	return &userRepositories{
		db: db,
	}
}
