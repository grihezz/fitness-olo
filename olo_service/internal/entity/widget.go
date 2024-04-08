package entity

type Widget struct {
	ID          int64  `db:"id"`
	Description string `db:"description"`
}
