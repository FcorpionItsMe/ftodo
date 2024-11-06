package service

import (
	"errors"
	"github.com/FcorpionItsMe/ftodo/internal/domain"
	"github.com/FcorpionItsMe/ftodo/internal/repository"
	autherrs "github.com/FcorpionItsMe/ftodo/internal/service/errors/auth"
	"github.com/FcorpionItsMe/ftodo/internal/utils/strs"
	"github.com/FcorpionItsMe/ftodo/pkg/hash/bcrypt"
	"github.com/FcorpionItsMe/ftodo/pkg/jwt_authorization"
	"log/slog"
)

type AuthService struct {
	repository repository.Repository
}

func NewAuthService(repository repository.Repository) *AuthService {
	return &AuthService{repository}
}

func (a AuthService) SignUp(user domain.SignUpUserInput) error {
	op := "service.auth.SignUp(): "
	isUserExist := a.repository.GetUserByLogin(user.Login) != nil
	if isUserExist {
		slog.Warn(op + "User already exists!")
		return autherrs.UserAlreadyExistError{}
	}
	if !strs.IsLower(user.Login) {
		slog.Warn(op + "Login is not in lower case!")
		return autherrs.IncorrectLoginError{AdditionalInfo: "Login is not in lower case!"}
	}
	hasher := bcrypt.New()
	hashedPassword, err := hasher.HashPassword(user.Pass, 12)
	if err != nil {
		slog.Warn(op + "Cannot hash password. " + err.Error())
		return autherrs.HashingError{}
	}
	user.Pass = hashedPassword
	err = a.repository.SaveUser(user)
	if err != nil {
		slog.Warn(op + "Cannot save user: " + err.Error())
		return err
	}
	slog.Info("User successfully sign-up!", slog.String("login", user.Login))
	return nil
}
func (a AuthService) SignIn(user domain.SignInUserInput) (string, error) {
	op := "service.auth.SignIn(): "
	if !strs.IsLower(user.Login) {
		slog.Warn(op + "Login is not in lower case!")
		return "", autherrs.IncorrectLoginError{AdditionalInfo: "Login is not in lower case!"}
	}
	//Check login
	if success, err := a.ValidUser(user.Login, user.Pass); !success || err != nil {
		slog.Warn(op + "Invalid username or password!")
		return "", err
	}
	jwtCrypter := jwt_authorization.NewJWTAuth()
	token, err := jwtCrypter.CreateToken(user)
	if err != nil {
		slog.Warn(op + "Cannot create token: " + err.Error())
		return "", err
	}
	return token, nil
}

func (a AuthService) ValidUser(login string, password string) (bool, error) {
	op := "service.auth.ValidUser(): "
	userInRep := a.repository.GetUserByLogin(login)
	if userInRep == nil {
		slog.Warn(op + "User does not exist!")
		return false, autherrs.UserDoesntExistError{Login: login}
	}

	hasher := bcrypt.New()
	if !hasher.ComparePasswordAndHash(password, userInRep.Pass) {
		slog.Warn(op + "User password does not match!")
		return false, errors.New("password does not match")
	}
	return true, nil
}
