package adapters

import (
	"context"
	"fmt"
	"somnium/internal/domain"
	"somnium/libs/postgres"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func (tsp *SomniumSystemPostgre) SignIn(ctx context.Context, input *domain.SignInRequest) (string, error) {
	const sql = `
		SELECT 
			u.id,
			u.hashed_password
		from "users" u
		where
			u.email = $1
	`
	var (
		userID     string
		hashedPass string
	)
	if err := tsp.pg.QueryRow(ctx, sql, input.Email).Scan(&userID, &hashedPass); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(input.Password)); err != nil {
		return "", err
	}
	return userID, nil
}

func (tsp *SomniumSystemPostgre) SignUp(ctx context.Context, input *domain.SignUpRequest) (string, error) {
	var userID string

	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		const sql = `
            INSERT INTO "users" (Email, hashed_password)
            VALUES ($1, $2)
            RETURNING id
        `

		if err := tx.QueryRow(ctx, sql, input.Email, string(hashedPassword)).Scan(&userID); err != nil {
			fmt.Println(err)
			return err
		}
		if err := tsp.assignRole(ctx, tx, userID, domain.RoleUser); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return userID, nil
}

func (tsp *SomniumSystemPostgre) AssignRole(ctx context.Context, userID, role string) error {
	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {
		if err := tsp.assignRole(ctx, tx, userID, role); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (tsp *SomniumSystemPostgre) assignRole(ctx context.Context, tx pgx.Tx, userID, role string) error {
	const sql = `
		insert into "user_role_bindings"
			(user_id, role_id)
		values
			(
				$1,
				(select id from user_roles where value = $2)
			)
	`
	if _, err := tx.Exec(ctx, sql, userID, role); err != nil {
		return err
	}

	return nil
}
