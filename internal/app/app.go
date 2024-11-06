package app

import (
	"github.com/FcorpionItsMe/ftodo/internal/config"
	"github.com/FcorpionItsMe/ftodo/internal/repository/postgres"
	"github.com/FcorpionItsMe/ftodo/internal/service"
	"github.com/FcorpionItsMe/ftodo/internal/transport/rest"
	pkgslog "github.com/FcorpionItsMe/ftodo/pkg/log/slog"
	"log"
	"log/slog"
	"net/http"
)

func Run(cfg *config.Config) {
	loggerCloseAction, err := pkgslog.NewLogger(cfg.Env)
	if err != nil {
		log.Fatal("Cannot setup logger! ", err)
		return
	}
	defer loggerCloseAction()
	slog.Info("App started!", slog.String("Env", cfg.Env))
	repository, err := postgres.New(cfg.DBConfig)
	if err != nil {
		slog.Error("Cannot connect repository! ", slog.String("Error", err.Error()))
		return
	}
	authService := service.NewAuthService(repository)
	//u := repository.GetUserByLogin("test_login1")
	//	Login:  "franchesco",
	//	Email:  "FranchescoBlas@gmail.com",
	//	Locale: "en",
	//	Pass:   "qwerty321",
	//})

	addres := ":" + cfg.ServerConfig.Port
	router := rest.New(authService)
	srver := http.Server{
		Addr:         addres,
		Handler:      router,
		ReadTimeout:  cfg.ServerConfig.Timeout,
		WriteTimeout: cfg.ServerConfig.Timeout,
		IdleTimeout:  cfg.ServerConfig.IdleTimeout,
	}
	slog.Info("Server on!")
	err = srver.ListenAndServe()
	if err != nil {
		slog.Error("Cannot start server! ", slog.String("Error", err.Error()))
	}
	slog.Info("Server off!")
	slog.Info("App stopped!")
}
