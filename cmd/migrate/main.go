package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/uptrace/bun"
    "github.com/uptrace/bun/dialect/pgdialect"
    "github.com/uptrace/bun/driver/pgdriver"
    "github.com/uptrace/bun/migrate"
)

func main() {
    dsn := "postgres://user:pass@localhost:5432/dbname?sslmode=disable"

    sqldb := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
    db := bun.NewDB(sqldb, pgdialect.New())

    migrator := migrate.NewMigrator(db, migrate.WithMigrationsDirectory("migrations"))

    ctx := context.Background()

    if err := migrator.Init(ctx); err != nil {
        log.Fatalf("erro ao inicializar migrator: %v", err)
    }

    // Executa comandos a partir da linha de comando
    if err := migrator.Run(ctx, os.Args...); err != nil {
        log.Fatalf("erro ao rodar migrations: %v", err)
    }

    fmt.Println("Migrations rodaram com sucesso ✅")
}
