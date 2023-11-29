package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type ExhangeRateRepository struct {
	db *sql.DB
}

func NewExhangeRateRepository(db *sql.DB) *ExhangeRateRepository {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS Exchange 
		(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, value FLOAT);
	`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return &ExhangeRateRepository{
		db: db,
	}
}

func (repo *ExhangeRateRepository) CreateExchange(exchangeValue float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		log.Println("Context error on func CreateExchange")
		return context.Canceled
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, "INSERT INTO Exchange (VALUE) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(exchangeValue)
	if err != nil {
		log.Println("Error inserting into the database on func CreateExchange")
		log.Println(err)
		return context.Canceled
	}
	select {
	case <-ctx.Done():
		log.Println("Context error on func CreateExchange")
		return context.Canceled
	default:
		return nil
	}
}
