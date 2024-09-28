package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

type Record struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

func NewStore(dataSource string) *Store {
	conn, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	return &Store{db: conn}
}

func (d *Store) SaveRecord(r Record) error {
	const op = "SaveRecord"
	q := "INSERT INTO urls VALUES (?, ?)"

	stmt, err := d.db.Prepare(q)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := stmt.Exec(r.Alias, r.URL); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Store) FindByURL(url string) (*Record, error) {
	const op = "FindByURL"
	q := "SELECT * FROM urls WHERE url = ?"

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var record Record

	if err := stmt.QueryRow(url).Scan(&record.Alias, &record.URL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &record, nil
}

func (s *Store) FindByAlias(alias string) (*Record, error) {
	const op = "FindByAlias"
	q := "SELECT * FROM urls WHERE alias = ?"

	stmt, err := s.db.Prepare(q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var record Record

	if err := stmt.QueryRow(alias).Scan(&record.Alias, &record.URL); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &record, nil
}
