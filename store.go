package main

import "database/sql"

type Store interface {
	CreateUser() error
}

type Storage struct {
	db *sql.DB
}

// Recebe Storage
func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser() error {
	return nil
}
