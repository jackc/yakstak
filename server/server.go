package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx"
	"github.com/jackc/yakstak/server/handlers"
)

type TodoRow struct {
	id   int32
	body string
	done bool
}

func Serve() {
	db, err := createDB()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Method("GET", "/", &handlers.YakstakIndex{DB: db})
	http.ListenAndServe(":3000", r)
}

func createDB() (*pgx.ConnPool, error) {
	connConfig := pgx.ConnConfig{Host: "/var/run/postgresql", Database: "yakstak_dev"}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 10,
	}

	return pgx.NewConnPool(poolConfig)
}
