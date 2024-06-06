goose -dir ./deploy/migration postgres "postgres://postgres:postgres@localhost:5432/somniumsystem?sslmode=disable" status

goose -dir ./deploy/migration postgres "postgres://postgres:postgres@localhost:5432/somniumsystem?sslmode=disable" up
