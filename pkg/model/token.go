package model

type TokenUser struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
	Role  string `db:"role"`
}
