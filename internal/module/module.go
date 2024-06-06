package module

import (
	"somnium/internal/adapters"
	"somnium/internal/domain"
	"somnium/libs/jwt"

	"somnium/internal/module/auth"
	"somnium/internal/module/job"
	"somnium/internal/module/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Model struct {
	*auth.AuthenticationModule
	*job.JobModule
	*user.UserModule
	db     domain.SomniumSystemPostgre
	Jwtcli domain.CredentialsRepository
}

func New(opts ...Option) *Model {
	deps := &dependencies{}
	deps.setDefaults()
	for _, opt := range opts {
		opt(deps)
	}
	svc := &Model{
		db:     adapters.NewSomniumSystemPostgre(deps.pg),
		Jwtcli: adapters.NewCredentialsRepositoryJWT(deps.jwtcli),
	}
	svc.AuthenticationModule = auth.NewAuthenticationModule(svc.db)
	svc.UserModule = user.NewUserModule(svc.db)
	svc.JobModule = job.NewJobModule(svc.db)
	return svc
}

type Option func(*dependencies)

type dependencies struct {
	pg     *pgxpool.Pool
	jwtcli *jwt.Client
}

func (d *dependencies) setDefaults() {
	// pass
}

func WithPostgres(pg *pgxpool.Pool) Option {
	return func(d *dependencies) {
		d.pg = pg
	}
}

func WithJWT(jwtcli *jwt.Client) Option {
	return func(d *dependencies) {
		d.jwtcli = jwtcli
	}
}
