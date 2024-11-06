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

func New(cfg config.DBConfig) (*Repository, error) {
	openedDB, err := sql.Open("postgres", //driver
		fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
			cfg.User,
			cfg.Pass,
			cfg.Host,
			cfg.DBName,
		),
	)
	if err != nil {
		slog.Error("Cannot connect to db!", slog.String("err", err.Error()))
		return nil, err
	}
	slog.Info("Connected to db!")
	return &Repository{db: openedDB}, nil
}
