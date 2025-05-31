package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"be.blog/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPG(cfg *configs.Config) *pgxpool.Pool {

	dsn := cfg.DBUrl

	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, dsn)

	dbpool.Config().MaxConnIdleTime = 1 * time.Minute
	dbpool.Config().MaxConnLifetime = 5 * time.Minute
	dbpool.Config().MaxConns = 10
	dbpool.Config().MinConns = 2

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	log.Println("Connected to database")
	return dbpool
}
