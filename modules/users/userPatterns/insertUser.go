package userPatterns

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/chon26909/e-commerce/modules/users"
	"github.com/jmoiron/sqlx"
)

type IInsertUser interface {
	Customer() (IInsertUser, error)
	Admin() (IInsertUser, error)
	Result() (*users.UserPassport, error)
}

type userRequest struct {
	id  string
	req *users.UserRegisterRequest
	db  *sqlx.DB
}

type customer struct {
	*userRequest
}

type admin struct {
	*userRequest
}

func InsertUser(db *sqlx.DB, req *users.UserRegisterRequest, isAdmin bool) IInsertUser {
	if isAdmin {
		return newAdmin(db, req)
	}
	return newCustomer(db, req)
}

func newCustomer(db *sqlx.DB, req *users.UserRegisterRequest) IInsertUser {
	return &customer{
		userRequest: &userRequest{
			req: req,
			db:  db,
		},
	}
}

func newAdmin(db *sqlx.DB, req *users.UserRegisterRequest) IInsertUser {
	return &admin{
		userRequest: &userRequest{
			req: req,
			db:  db,
		},
	}
}

func (f *userRequest) Customer() (IInsertUser, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
		INSERT INTI "users" (
			"email",
			"password",
			"username",
			"role_id"
		)
		VALUES 
			($1, $2, $3, $4)
		RETURNING "id";
	`

	if err := f.db.QueryRowContext(
		ctx,
		query,
		f.req.Email,
		f.req.Password,
		f.req.Username,
	); err != nil {
		return nil, fmt.Errorf("insert user failed: %v", err)
	}

	return nil, nil
}

func (f *userRequest) Admin() (IInsertUser, error) {
	return nil, nil
}

// Result implements IInsertUser.
func (f *userRequest) Result() (*users.UserPassport, error) {
	query := `
		SELECT 
			json_build_object(
				'user',"t"
				'token, NULL
			)
		FROM (
			SELECT 
				"u"."id",
				"u"."email",
				"u"."username",
				"u"."role_id"
			FROM "users" "u"
			WHERE "u"."id" = $1
		) AS "t"
	`

	data := make([]byte, 0)

	if err := f.db.Get(&data, query, f.id); err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}

	user := new(users.UserPassport)
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}

	return user, nil
}
