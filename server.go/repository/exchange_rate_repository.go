package repository

import (
	"database/sql"
	"log"
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
	defer stmt.Close()
	_, err = stmt.Exec(exchangeValue)
	if err != nil {
		return err
	}
	return nil
}
