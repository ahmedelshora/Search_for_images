package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	conn *sql.DB
}

type Product struct {
	ID   int
	Name string
}

func NewDatabase(cfg *Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &Database{conn: db}, nil
}

func (db *Database) GetProducts() ([]Product, error) {
	rows, err := db.conn.Query(`
		SELECT product_id, name
		FROM oc_product_description
		WHERE language_id = 1
	`)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return products, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}
