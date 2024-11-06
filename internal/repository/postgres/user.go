package postgres

import (
	"fmt"
	"github.com/FcorpionItsMe/ftodo/internal/domain"
	"github.com/FcorpionItsMe/ftodo/internal/utils/repository/postgres/parser"
	"github.com/FcorpionItsMe/ftodo/internal/utils/repository/postgres/pq_key"
	"log/slog"
)

func (r Repository) SaveUser(inputUser domain.SignUpUserInput) error {
	op := "Repository.SaveUser(): "
	table := pq_key.UserTableKeys
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES($1,$2,$3,$4)",
		table.Name,
		table.Columns.Email,
		table.Columns.Login,
		table.Columns.Pass,
		table.Columns.Locale,
	)
	preparedQuery, err := r.db.Prepare(query)
	if err != nil {
		slog.Warn(op + "Cannot prepare query to execute!")
		return err
	}
	_, err = preparedQuery.Exec(inputUser.Email, inputUser.Login, inputUser.Pass, inputUser.Locale)
	if err != nil {
		slog.Warn(op + "Error while executing query to insert into User table!")
		return err
	}
	return nil
}

func (r Repository) GetUserByLogin(login string) *domain.User {
	//op := "Repository.GetUserByLogin(): "
	table := pq_key.UserTableKeys
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table.Name, table.Columns.Login)
	row := r.db.QueryRow(query, login)
	userFromRow, err := parser.ParseUserFromRow(row)
	if err != nil {
		return nil
	}
	return userFromRow
}
