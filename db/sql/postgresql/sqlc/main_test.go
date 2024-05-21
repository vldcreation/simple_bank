package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var (
		err               error
		POSTGRES_USER     = os.Getenv("POSTGRES_USER")
		POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		POSTGRES_DB       = os.Getenv("POSTGRES_DB")
		POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
		POSTGRES_PORT     = os.Getenv("POSTGRES_PORT")
	)

	dbSource := "postgresql://" + POSTGRES_USER + ":" + POSTGRES_PASSWORD + "@" + POSTGRES_HOST + ":" + POSTGRES_PORT + "/" + POSTGRES_DB + "?sslmode=disable"

	testDB, err = sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
