package test_utils

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func InitTxDbDatabase(t *testing.T) (*sql.DB, error) {
	t.Helper()
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("error loading .env file")
	}
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TEST"),
	)

	txdb.Register("txdb", "mysql", dataSource)
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
