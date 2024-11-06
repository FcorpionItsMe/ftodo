package handler

import (
	"encoding/json"
	"fmt"
	"github.com/FcorpionItsMe/ftodo/internal/domain"
	"github.com/FcorpionItsMe/ftodo/internal/service"
	"github.com/FcorpionItsMe/ftodo/internal/utils/rest"
	"log/slog"
	"net/http"
)

const (
	AUTH_SIGN_UP_ROUTE = "POST /authorization/sign-up/"
	AUTH_SIGN_IN_ROUTE = "POST /authorization/sign-in/"
)

type AuthHandlerGroup struct {
	service *service.AuthService
}

func NewAuthHandlersGroup(service *service.AuthService) *AuthHandlerGroup {
	return &AuthHandlerGroup{service: service}
}

func (g AuthHandlerGroup) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("method", r.Method),
		slog.String("url", r.URL.String()),
	)
	w.Header().Set("Content-Type", "application/json")
	if r.ContentLength == 0 || r.Header.Get("Content-Type") != "application/json" {
		logger.Warn("Request with empty content length or data is not json!")
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Body is empty or not json")); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	body, err := rest.ReadBody(r)
	if err != nil {
		logger.Warn("Error reading body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte("Wrong json!")); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	userInfo := domain.SignUpUserInput{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		logger.Warn("Error unmarshalling body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte("Wrong json!")); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	err = g.service.SignUp(userInfo)
	if err != nil {
		logger.Warn("Error signing up", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(fmt.Sprintf("{\"message\": \"%s\"}", err))); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"You signed up successfully!\"}"))
}
func (g AuthHandlerGroup) SignInHandler(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("method", r.Method),
		slog.String("url", r.URL.String()),
	)
	w.Header().Set("Content-Type", "application/json")
	if r.ContentLength == 0 || r.Header.Get("Content-Type") != "application/json" {
		logger.Warn("Request with empty content length or data is not json!")
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Body is empty or not json")); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	body, err := rest.ReadBody(r)
	if err != nil {
		logger.Warn("Error reading body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte("Wrong json!")); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	userInfo := domain.SignInUserInput{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		logger.Warn("Error unmarshalling body", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte("Wrong json!")); err != nil {
			logger.Warn("Failed to write response")
		}
		return
	}
	token, err := g.service.SignIn(userInfo)
	if err != nil {
		logger.Warn("Error signing in", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(fmt.Sprintf("{\"message\": \"%s\"}", err))); err != nil {
			slog.Error("Failed to write response")
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf("{\"token\": \"%s\"}", token)))
	if err != nil {
		slog.Error("Failed to write response")
	}
}
