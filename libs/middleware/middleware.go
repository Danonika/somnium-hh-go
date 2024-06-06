package mw

import (
	"somnium/internal/adapters"
	"somnium/internal/domain"
	"somnium/libs/jwt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Middleware struct {
	jc *jwt.Client
	db domain.MiddlewareRepository
}

func NewMiddleware(jc *jwt.Client, pg *pgxpool.Pool) *Middleware {
	return &Middleware{
		jc: jc,
		db: adapters.NewMiddlewarePostgre(pg),
	}
}
