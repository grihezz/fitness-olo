package entity

type Widget struct {
	ID   int64  `db:"id"`
	Data string `db:"data"`
}
