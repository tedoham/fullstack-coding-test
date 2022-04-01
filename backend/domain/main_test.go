package domain

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/products?sslmode=disable"
)

// var testQueries *sqlx.DB

func TestMain(m *testing.M) {
	_, err := sqlx.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// testQueries = sqlx.

	os.Exit(m.Run())
}
