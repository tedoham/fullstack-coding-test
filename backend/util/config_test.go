package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	require := require.New(t)

	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_SOURCE", "postgres://root:secret@localhost:5432/products?sslmode=disable")
	os.Setenv("SERVER_ADDRESS", "localhost:8080")

	actual, err := LoadConfig("../")
	require.NoError(err)

	expected := Config{
		DB_DRIVER:      "postgres",
		DB_SOURCE:      "postgres://root:secret@localhost:5432/products?sslmode=disable",
		SERVER_ADDRESS: "localhost:8080",
	}
	require.Equal(expected, actual)
}
