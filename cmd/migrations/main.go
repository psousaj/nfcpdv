package main

import (
	"context"
	"fmt"
	_ "log"
	_ "os"

	_ "github.com/uptrace/bun/migrate"

	"github.com/psousaj/nfcpdv/config"
	"github.com/psousaj/nfcpdv/internal/domain"
	"github.com/psousaj/nfcpdv/internal/infra/db"
)

func main() {
	cfg := config.LoadConfig()
	db := db.ConnectPostgres(cfg.DBUrl)
	ctx := context.Background()

	// Run CREATE TABLE statements
	_, job_err := db.NewCreateTable().
    Model((*domain.Job)(nil)).
    IfNotExists().
    Exec(ctx)

    if job_err != nil {
        panic(job_err)
    }

	fmt.Println("Migrations rodaram com sucesso ✅")
}
