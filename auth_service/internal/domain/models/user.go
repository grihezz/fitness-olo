package models

type User struct {
	ID           int64  `db:"id"`
	Email        string `db:"email"`
	Role         string `db:"role"`
	PassHash     []byte `db:"password_hash"`
	DateRegister string `db:"date_register"`
}
