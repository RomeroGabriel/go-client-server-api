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
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return &ExhangeRateRepository{
		db: db,
	}
}

func (repo *ExhangeRateRepository) CreateExchange(exchangeValue float64) error {
	stmt, err := repo.db.Prepare("INSERT INTO Exchange (VALUE) VALUES (?)")
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	defer stmt.Close()
	_, err = stmt.Exec(exchangeValue)
	if err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		log.Println("Request Timeout")
		log.Println("Request Timeout inserting into the database, func CreateExchange")
		return context.Canceled
	default:
		return nil
	}
}
