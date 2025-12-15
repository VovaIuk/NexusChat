package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(tag, username, password string) (User, error) {
	if tag == "" {
		return User{}, errors.New("tag cannot be empty")
	}
	if username == "" {
		return User{}, errors.New("username cannot be empty")
	}
	if len(password) < 8 {
		return User{}, errors.New("password must be at least 8 characters long")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return User{}, err
	}

	return User{
		Tag:      tag,
		Username: username,
		Password: string(hashedPassword),
	}, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	)
	return err == nil
}
