package adapters

import (
	"context"
	"somnium/internal/domain"
	"somnium/libs/postgres"

	"github.com/jackc/pgx/v4"
)

func (tsp *SomniumSystemPostgre) GetRoles(ctx context.Context, userID string) ([]string, error) {
	const sql = `
		select
			ur.value
		from user_role_bindings urb
		join user_roles ur on ur.id = urb.role_id
		where
			urb.user_id = $1
	`

	var (
		userRoles []string
	)

	rows, err := tsp.pg.Query(ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userRole string
		if err := rows.Scan(&userRole); err != nil {
			return nil, err
		}
		userRoles = append(userRoles, userRole)
	}

	return userRoles, nil
}

func (tsp *SomniumSystemPostgre) UpdateUser(ctx context.Context, input *domain.UpdateUserRequest) error {
	return postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {
		const sql = `
            UPDATE "users" 
            SET  
                FIO = $1, 
                Contact = $2, 
                resume_title = $3, 
                resume_link = $4, 
                skills = $5
            WHERE ID = $6
        `
		_, err := tx.Exec(ctx, sql, input.UserInfo.Name, input.UserInfo.ContactNum, input.UserInfo.ResumeTitle, input.UserInfo.ResumeLink, input.UserInfo.Skills, input.UserID)
		return err
	})
}

func (tsp *SomniumSystemPostgre) GetUser(ctx context.Context, input *domain.GetUserRequest) (*domain.UserInfo, error) {
	const sql = `
        SELECT 
            Email,
            FIO,
            Contact,
            resume_title,
            resume_link,
            skills
        FROM "users"
        WHERE ID = $1
    `
	var userInfo domain.UserInfo
	err := tsp.pg.QueryRow(ctx, sql, input.UserID).Scan(&userInfo.Email, &userInfo.Name, &userInfo.ContactNum, &userInfo.ResumeTitle, &userInfo.ResumeLink, &userInfo.Skills)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
