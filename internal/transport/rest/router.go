package rest

import (
	"github.com/FcorpionItsMe/ftodo/internal/service"
	"github.com/FcorpionItsMe/ftodo/internal/transport/rest/handler"
	"net/http"
)

type Router struct {
	*http.ServeMux
}

func New(authService *service.AuthService) *Router {
	muxer := &http.ServeMux{}
	muxer.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Main PAGE!"))
	})
	auth := handler.NewAuthHandlersGroup(authService)
	muxer.HandleFunc(handler.AUTH_SIGN_UP_ROUTE, auth.SignUpHandler)
	muxer.HandleFunc(handler.AUTH_SIGN_IN_ROUTE, auth.SignInHandler)
	return &Router{muxer}
}
