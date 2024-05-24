package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/vldcreation/simple_bank/api"
	"github.com/vldcreation/simple_bank/app"
	db "github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc"
)

const (
	dbdriver = "postgres"
)

func main() {
	var (
		err error
		cfg = app.NewConfigFromEnv("")
	)

	dbSource := "postgresql://" + cfg.DB.User + ":" + cfg.DB.Password + "@" + cfg.DB.Host + ":" + cfg.DB.Port + "/" + cfg.DB.Database + "?sslmode=disable"

	conn, err := sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(":8080")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
