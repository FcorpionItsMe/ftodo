package pq_key

import "github.com/lib/pq"

var UserTableKeys = userTable{
	Name: pq.QuoteIdentifier("user"),
	Columns: userTableColumns{
		Id:     pq.QuoteIdentifier("id"),
		Email:  pq.QuoteIdentifier("email"),
		Login:  pq.QuoteIdentifier("login"),
		Pass:   pq.QuoteIdentifier("password"),
		Locale: pq.QuoteIdentifier("locale"),
		Date:   pq.QuoteIdentifier("date"),
	},
}

type userTable struct {
	Name    string
	Columns userTableColumns
}
type userTableColumns struct {
	Id     string
	Email  string
	Login  string
	Pass   string
	Locale string
	Date   string
}
