package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/berci_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Could not connect to the DB", err)
	}
	defer pool.Close()

	testQueries = New(pool)

	os.Exit(m.Run())
}
