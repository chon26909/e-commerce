package userRepositories

import (
	"context"
	"fmt"
	"time"

	"github.com/chon26909/e-commerce/modules/users"
	"github.com/chon26909/e-commerce/modules/users/userPatterns"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	InsertUser(req *users.UserRegisterRequest, isAdmin bool) (*users.UserPassport, error)
	FindOneUserbyEmail(email string) (*users.UserCredentialCheck, error)
	InsertOAuth(req *users.UserPassport) error
	FindOneOauth(refreshToken string) (*users.Oauth, error)
	UpdateOauth(req *users.UserToken) error
	GetProfile(userId string) (*users.User, error)
	DeleteOauth(oauthId string) error
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
			"role_id"
		FROM "users"
		WHERE "email" = $1;
	`

	user := new(users.UserCredentialCheck)
	if err := r.db.Get(user, query, email); err != nil {
		return nil, fmt.Errorf("user not found : %v", err)
	}

	return user, nil
}

func (r *userRepository) InsertOAuth(req *users.UserPassport) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	query := `
		INSERT INTO "oauth" (
			"user_id",
			"refresh_token",
			"access_token"
		) 
		VALUES ($1, $2, $3)
		RETURNING "id";
	`

	if err := r.db.QueryRowContext(
		ctx,
		query,
		req.User.Id,
		req.Token.RefreshToken,
		req.Token.AccessToken,
	).Scan(&req.Token.Id); err != nil {
		return fmt.Errorf("insert oauth failed: %v", err)
	}

	return nil
}

func (r *userRepository) FindOneOauth(refreshToken string) (*users.Oauth, error) {

	query := `
		SELECT 
			"id",
			"user_id"
		FROM "oauth"
		WHERE "refresh_token" = $1;
	`

	oauth := new(users.Oauth)
	if err := r.db.Get(oauth, query, refreshToken); err != nil {
		return nil, fmt.Errorf("oauth not found: %v", err)
	}
	return oauth, nil
}

func (r *userRepository) UpdateOauth(req *users.UserToken) error {

	query := `
		UPDATE 	"oauth" 
		SET 	"access_token" = :access_token,
				"refresh_token" = :refresh_token
		WHERE "id" = :id;
	`
	if _, err := r.db.NamedExecContext(context.Background(), query, req); err != nil {
		return fmt.Errorf("update oauth failed: %v", err)
	}

	return nil
}

func (r userRepository) GetProfile(userId string) (*users.User, error) {

	query := ` 
		SELECT
        	"id",
        	"email",
        	"username",
        	"role_id"
    	FROM "users"
    	WHERE "id" = $1;`

	profile := new(users.User)

	if err := r.db.Get(profile, query, userId); err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}
	return profile, nil
}

func (r *userRepository) DeleteOauth(oauthId string) error {
	query := `DELETE FROM "oauth" WHERE "id" = $1;`

	if _, err := r.db.ExecContext(context.Background(), query, oauthId); err != nil {
		return fmt.Errorf("oauth not found")
	}

	return nil
}
