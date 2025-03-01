package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/vldcreation/simple_bank/api"
	"github.com/vldcreation/simple_bank/app"
	db "github.com/vldcreation/simple_bank/db/sql/postgresql/sqlc"
	"github.com/vldcreation/simple_bank/token"
)

const (
	dbdriver = "postgres"
)

func main() {
	var (
		err error
		cfg = app.NewConfigFromYaml("")
	)

	dbSource := "postgresql://" + cfg.DB.User + ":" + cfg.DB.Password + "@" + cfg.DB.Host + ":" + cfg.DB.Port + "/" + cfg.DB.Database + "?sslmode=disable"

	conn, err := sql.Open(dbdriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	tokenMaker, err := token.NewPasetoMaker(cfg.Token.SecretKey)
	if err != nil {
		log.Fatalf("cannot init token maker: %+v\n", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store, tokenMaker, cfg)

	err = server.Start(fmt.Sprintf(":%s", cfg.APP.Port))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
