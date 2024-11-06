package parser

import (
	"database/sql"
	"github.com/FcorpionItsMe/ftodo/internal/domain"
)

func ParseUserFromRow(row *sql.Row) (*domain.User, error) {
	user := &domain.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Login, &user.Pass, &user.Locale, &user.Date)
	if err != nil {
		return nil, err
	}
	return user, nil
}
