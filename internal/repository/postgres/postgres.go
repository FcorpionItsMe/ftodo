package postgres

import (
	"database/sql"
	"fmt"
	"github.com/FcorpionItsMe/ftodo/internal/config"
	_ "github.com/lib/pq"
	"log/slog"
)

type Repository struct {
	db *sql.DB
}

func New(cfg config.DBConfig, logger *slog.Logger) (*Repository, error) {
	openedDB, err := sql.Open("postgres", //driver
		fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%t",
			cfg.User,
			cfg.Pass,
			cfg.Host,
			cfg.DBName,
			cfg.SSLMode,
		),
	)
	if err != nil {
		logger.Error("Cannot connect to db!", slog.String("err", err.Error()))
		return nil, err
	}
	logger.Info("Connected to db!")
	return &Repository{db: openedDB}, nil
}
