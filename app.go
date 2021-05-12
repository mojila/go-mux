package main

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("host=%s port=%d "+
		"user=%s password=%s dbname=%s sslmode=%s",
		"backbone.machinevision.global", 5432, user, password, dbname, "disable")
	var err error

	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}
func (a *App) Run(addr string) {}
