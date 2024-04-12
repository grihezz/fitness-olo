package entity

type Article struct {
	ID     int64  `db:"id"`
	Header string `db:"header"`
	Body   string `db:"body"`
}
