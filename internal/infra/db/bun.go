package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewBunDB(dataSourceName string) *bun.DB {

	sqldb, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
