package userPatterns

import "github.com/chon26909/e-commerce/modules/users"

type IInsertUser interface {
	Customer() (IInsertUser, error)
	Admin() (IInsertUser, error)
	Result(*users.User) error
}
