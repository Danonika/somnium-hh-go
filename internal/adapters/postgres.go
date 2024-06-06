package adapters

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type SomniumSystemPostgre struct {
	pg *pgxpool.Pool
}

func NewSomniumSystemPostgre(pg *pgxpool.Pool) *SomniumSystemPostgre {
	return &SomniumSystemPostgre{
		pg: pg,
	}
}

type MiddlewarePostgre struct {
	pg *pgxpool.Pool
}

func NewMiddlewarePostgre(pg *pgxpool.Pool) *MiddlewarePostgre {
	return &MiddlewarePostgre{
		pg: pg,
	}
}
