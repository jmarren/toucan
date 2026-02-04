package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Dbtx  *pgxpool.Pool
	Query *Queries
)

func Init(ctx context.Context) error {
	connStr := os.Getenv("DB_URL")
	fmt.Printf("connection string: %s\n", connStr)

	DbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Println("error creating pool")
		return err
	}

	Query = New(DbPool)

	if Query == nil {
		return fmt.Errorf("error: nil db.Query !")
	}

	fmt.Printf("database initialized successfully\n")
	return nil
}
