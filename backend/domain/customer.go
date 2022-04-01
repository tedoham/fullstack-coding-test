package domain

type Customer struct {
	ID            int    `db:"id"`
	Name          string `db:"name"`
	ContactNumber string `db:"contact_number"`
}
