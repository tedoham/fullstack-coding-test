package domain

type DeliveryStatus struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
