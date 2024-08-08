package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := &mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBUser,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(cfg)
	db, err := sqlStorage.Init() // Retorna db pront -> com tabelas e colunos necessárias
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	api := NewAPIServer(":8080", store)
	api.Serve()
}
