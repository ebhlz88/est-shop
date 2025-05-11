package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=user dbname=mydatabase password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) InitTables() error {
	if err := s.CreateUserTable(); err != nil {
		return err
	}
	if err := s.CreateProductTable(); err != nil {
		return err
	}
	if err := s.CreateOrderTable(); err != nil {
		return err
	}
	if err := s.CreateProfitTable(); err != nil {
		return err
	}
	if err := s.CreateStockTable(); err != nil {
		return err
	}
	if err := s.CreateInvestorTable(); err != nil {
		return err
	}
	return nil
}
