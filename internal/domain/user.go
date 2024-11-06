package domain

import (
	"log/slog"
	"strconv"
	"time"
)

type User struct {
	Id     int       `json:"id" validate:"required"`
	Email  string    `json:"email" validate:"required,email"`
	Login  string    `json:"login" validate:"required,gte=3,lte=16"`
	Pass   string    `json:"pass" validate:"required,gte=8,lte=512"`
	Locale string    `json:"locale" validate:"required,gte=2, lte=5"`
	Date   time.Time `json:"date" validate:"required"`
}

func NewUser(id int, email, name, pass, locale string, date time.Time) *User {
	createdUser := &User{id, email, name, pass, locale, date}
	err := Validator.Struct(createdUser)
	if err != nil {
		slog.Warn("Cannot validate user",
			slog.String("UserId", strconv.Itoa(id)),
			slog.String("Error", err.Error()),
		)
		return nil
	}
	return createdUser
}

type SignUpUserInput struct {
	Email  string `json:"email" validate:"required,email"`
	Login  string `json:"login" validate:"required,gte=3,lte=16"`
	Pass   string `json:"pass" validate:"required,gte=8,lte=512"`
	Locale string `json:"locale" validate:"required,gte=2, lte=5"`
}

func NewSignUpUserInput(email, name, pass, locale string) *SignUpUserInput {
	createdInputUser := &SignUpUserInput{email, name, pass, locale}
	err := Validator.Struct(createdInputUser)
	if err != nil {
		slog.Warn("Cannot validate inputUser", slog.String("Error", err.Error()))
		return nil
	}
	return createdInputUser
}

type SignInUserInput struct {
	Login string `json:"login" validate:"required,gte=3,lte=16"`
	Pass  string `json:"pass" validate:"required,gte=8,lte=512"`
}

func NewSignInUserInput(name, pass string) *SignInUserInput {
	createdInputUser := &SignInUserInput{name, pass}
	err := Validator.Struct(createdInputUser)
	if err != nil {
		slog.Warn("Cannot validate inputUser", slog.String("Error", err.Error()))
		return nil
	}
	return createdInputUser
}
