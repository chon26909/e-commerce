package users

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	UserName string `db:"user_name" json:"user_name"`
	RoleId   int    `db:"role_id" json:"role_id"`
}

type UserRegisterRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Username string `db:"username" json:"username"`
}

type UserCredential struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserCredentialCheck struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Username string `db:"username"`
	RoleId   int    `db:"role_id"`
}

func (obj *UserRegisterRequest) BcryptHash() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), 10)
	if err != nil {
		return fmt.Errorf("hashed password failed: %v", err)
	}
	obj.Password = string(hashedPassword)

	return nil
}

func (obj *UserRegisterRequest) IsEmail() bool {
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, obj.Email)
	if err != nil {
		return false
	}

	return match
}

type UserPassport struct {
	User  *User   `json:"user"`
	Token *string `json:"token"`
}

type UserToken struct {
	Id           string `db:"id" json:"id"`
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}

type UserClaims struct {
	Id     string `db:"id" json:"id"`
	RoleId int    `db:"role" json:"role"`
}
