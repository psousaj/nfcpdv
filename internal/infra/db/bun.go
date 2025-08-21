package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func ConnectPostgres(dataSourceName string) *bun.DB {

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dataSourceName)))
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}
