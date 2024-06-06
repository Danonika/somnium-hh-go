package job

import (
	"context"
	"somnium/internal/domain"
)

type JobModule struct {
	db domain.SomniumSystemPostgre
}

func NewJobModule(db domain.SomniumSystemPostgre) *JobModule {
	return &JobModule{
		db: db,
	}
}

func (jm *JobModule) AddSkill(ctx context.Context, req *domain.AddSkillRequest) error {
	return jm.db.AddSkill(ctx, req.Skill)
}

func (jm *JobModule) SkillPool(ctx context.Context) ([]string, error) {
	return jm.db.SkillPool(ctx)
}

func (jm *JobModule) GetJob(ctx context.Context, req *domain.GetJobRequest) (*domain.JobInfo, error) {
	return jm.db.GetJob(ctx, req)
}

func (jm *JobModule) AddJob(ctx context.Context, req *domain.AddJobRequest) (string, error) {
	return jm.db.AddJob(ctx, req)
}

func (jm *JobModule) DeleteJob(ctx context.Context, req *domain.DeleteJobRequest) error {
	return jm.db.DeleteJob(ctx, req)
}

func (jm *JobModule) UpdateJob(ctx context.Context, req *domain.UpdateJobRequest) error {
	return jm.db.UpdateJob(ctx, req)
}

func (jm *JobModule) ListJobs(ctx context.Context) ([]domain.JobInfo, error) {
	return jm.db.ListJobs(ctx)
}
