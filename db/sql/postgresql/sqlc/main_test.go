package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vldcreation/simple_bank/app"
)

const (
	dbdriver = "postgres"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var (
		err error
		cfg = app.NewConfigFromEnv("../../../../config")
	)

	dbSource := "postgresql://" + cfg.DB.User + ":" + cfg.DB.Password + "@" + cfg.DB.Host + ":" + cfg.DB.Port + "/" + cfg.DB.Database + "?sslmode=disable"

	testDB, err = sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
