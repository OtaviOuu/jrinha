package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg *mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Garantir que a conexão existe
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Conectado com MySQL")

	return &MySQLStorage{db: db}
}

// Inicializar tabelas / colunas
func (s *MySQLStorage) Init() (*sql.DB, error) {
	return s.db, nil
}
