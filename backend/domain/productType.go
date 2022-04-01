package domain

type ProductType struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
