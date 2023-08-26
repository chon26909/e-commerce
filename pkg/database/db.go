package database

import (
	"log"

	"github.com/chon26909/e-commerce/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(config config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", config.Url())
	if err != nil {
		log.Fatalf("connect db failed %v\n", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConnections())

	return db
}
