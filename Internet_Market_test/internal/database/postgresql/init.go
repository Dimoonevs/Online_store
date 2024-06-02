package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"time"
)

var countsTryConnect int32

// ConnectToPostgres connect to db postgresql
func ConnectToPostgres() *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	for {
		pool, err := pgxpool.Connect(context.Background(), dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			countsTryConnect++
		} else {
			log.Println("Connected to Postgres!")
			return pool
		}

		if countsTryConnect > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
