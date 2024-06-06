package adapters

import (
	"context"
	"somnium/internal/domain"
	"somnium/libs/postgres"

	"github.com/jackc/pgx/v4"
)

func (tsp *SomniumSystemPostgre) AddSkill(ctx context.Context, skill string) error {
	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {

		const sql = `
		INSERT INTO skills (skill)
		VALUES ($1)
	`
		_, err := tx.Exec(ctx, sql, skill)
		return err
	})
	return err
}

func (tsp *SomniumSystemPostgre) SkillPool(ctx context.Context) ([]string, error) {
	const sql = `
		SELECT skill
		FROM skills
	`
	rows, err := tsp.pg.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []string
	for rows.Next() {
		var skill string
		if err := rows.Scan(&skill); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (tsp *SomniumSystemPostgre) AddJob(ctx context.Context, req *domain.AddJobRequest) (string, error) {

	var jobID string
	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {
		const sql = `
		INSERT INTO job (title, salary, categoryid, description, location, contact_number, contact_email, skills, expire_date, statusid)
		VALUES ($1, $2, (SELECT id FROM category WHERE name = $3), $4, $5, $6, $7, $8, $9, (SELECT id FROM status WHERE name = $10))
		RETURNING id
	`
		row := tx.QueryRow(ctx, sql, req.Title, req.Salary, req.Category, req.Descr, req.Location, req.ContactNum, req.ContactEmai, req.Skills, req.Expiry, req.Status)
		err := row.Scan(&jobID)
		return err
	})
	return jobID, err
}

func (tsp *SomniumSystemPostgre) GetJob(ctx context.Context, req *domain.GetJobRequest) (*domain.JobInfo, error) {
	const sql = `
		SELECT
			j.id, j.title, j.salary, c.name as category, j.description, j.location, j.contact_number, j.contact_email, j.skills,
			j.date_posted, j.expire_date, s.name as status, j.applicants_count
		FROM job j
		INNER JOIN category c ON j.categoryid = c.id
		INNER JOIN status s ON j.statusid = s.id
		WHERE j.id = $1
	`
	var jobInfo domain.JobInfo
	err := tsp.pg.QueryRow(ctx, sql, req.JobID).Scan(&jobInfo.JobID, &jobInfo.Title, &jobInfo.Salary, &jobInfo.Category, &jobInfo.Descr,
		&jobInfo.Location, &jobInfo.ContactNum, &jobInfo.ContactEmail, &jobInfo.Skills, &jobInfo.DatePosted, &jobInfo.Expiry,
		&jobInfo.Status, &jobInfo.ApplicantsCount)
	if err != nil {
		return nil, err
	}
	return &jobInfo, nil
}

func (tsp *SomniumSystemPostgre) DeleteJob(ctx context.Context, req *domain.DeleteJobRequest) error {
	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {

		const sql = `
		DELETE FROM job WHERE id = $1
	`
		_, err := tx.Exec(ctx, sql, req.JobID)
		return err
	})
	return err
}

func (tsp *SomniumSystemPostgre) UpdateJob(ctx context.Context, req *domain.UpdateJobRequest) error {
	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {

		const sql = `
		UPDATE job
		SET title = $1, salary = $2, categoryid = (SELECT id FROM category WHERE name = $3),
			description = $4, location = $5, contact_number = $6, contact_email = $7,
			skills = $8, expire_date = $9, statusid = (SELECT id FROM status WHERE name = $10)
		WHERE id = $11
	`
		_, err := tx.Exec(ctx, sql, req.JobInfo.Title, req.JobInfo.Salary, req.JobInfo.Category, req.JobInfo.Descr, req.JobInfo.Location,
			req.JobInfo.ContactNum, req.JobInfo.ContactEmail, req.JobInfo.Skills, req.JobInfo.Expiry, req.JobInfo.Status, req.JobID)
		return err
	})
	return err
}

func (tsp *SomniumSystemPostgre) ListJobs(ctx context.Context) ([]domain.JobInfo, error) {
	const sql = `
		SELECT
			j.id, j.title, j.salary, c.name as category, j.description, j.location, j.contact_number, j.contact_email, j.skills,
			j.date_posted, j.expire_date, s.name as status, j.applicants_count
		FROM job j
		INNER JOIN category c ON j.categoryid = c.id
		INNER JOIN status s ON j.statusid = s.id
	`
	rows, err := tsp.pg.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []domain.JobInfo
	for rows.Next() {
		var jobInfo domain.JobInfo
		err := rows.Scan(&jobInfo.JobID, &jobInfo.Title, &jobInfo.Salary, &jobInfo.Category, &jobInfo.Descr, &jobInfo.Location,
			&jobInfo.ContactNum, &jobInfo.ContactEmail, &jobInfo.Skills, &jobInfo.DatePosted, &jobInfo.Expiry, &jobInfo.Status,
			&jobInfo.ApplicantsCount)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, jobInfo)
	}
	return jobs, nil
}
func (tsp *SomniumSystemPostgre) SwitchJobStatus(ctx context.Context, req *domain.JobSwitcherRequest) error {

	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {
		const sql = `
			UPDATE job
			SET statusid = (
				SELECT id FROM status WHERE name = (
					CASE 
						WHEN (SELECT name FROM status WHERE id = job.statusid) = 'Active' THEN 'Inactive' 
						ELSE 'Active'
					END
				)
			)
			WHERE id = $1
		`
		_, err := tx.Exec(ctx, sql, req.JobID)
		return err
	})
	return err
}

func (tsp *SomniumSystemPostgre) ApplyJob(ctx context.Context, req *domain.ApplyRequest) error {
	err := postgres.InTransaction(ctx, tsp.pg, func(ctx context.Context, tx pgx.Tx) error {

		const sql = `
		INSERT INTO job_applications (user_id, job_id)
		VALUES ($1, $2)
	`
		_, err := tx.Exec(ctx, sql, req.UserID, req.JobID)
		return err
	})
	return err
}

func (tsp *SomniumSystemPostgre) GetAppliedJobs(ctx context.Context, req *domain.ApplyHistoryRequest) ([]domain.JobInfo, error) {
	const sql = `
		SELECT
			j.id, j.title, j.salary, c.name as category, j.description, j.location, j.contact_number, j.contact_email, j.skills,
			j.date_posted, j.expire_date, s.name as status, j.applicants_count
		FROM job j
		INNER JOIN job_applications ja ON j.id = ja.job_id
		INNER JOIN category c ON j.categoryid = c.id
		INNER JOIN status s ON j.statusid = s.id
		WHERE ja.user_id = $1
	`
	rows, err := tsp.pg.Query(ctx, sql, req.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []domain.JobInfo
	for rows.Next() {
		var jobInfo domain.JobInfo
		err := rows.Scan(&jobInfo.JobID, &jobInfo.Title, &jobInfo.Salary, &jobInfo.Category, &jobInfo.Descr, &jobInfo.Location,
			&jobInfo.ContactNum, &jobInfo.ContactEmail, &jobInfo.Skills, &jobInfo.DatePosted, &jobInfo.Expiry, &jobInfo.Status,
			&jobInfo.ApplicantsCount)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, jobInfo)
	}
	return jobs, nil
}
