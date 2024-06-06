package somniumsystem

import (
	"context"
	"somnium/internal/domain"
	desc "somnium/pkg/api/somnium/v1"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) AddSkill(ctx context.Context, req *desc.AddSkillRequest) (*emptypb.Empty, error) {
	err := s.module.AddSkill(ctx, &domain.AddSkillRequest{
		Skill: req.Skill,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *service) SkillPool(ctx context.Context, _ *emptypb.Empty) (*desc.SkillPoolResponse, error) {
	skills, err := s.module.SkillPool(ctx)
	if err != nil {
		return nil, err
	}
	return &desc.SkillPoolResponse{Skills: skills}, nil
}

func (s *service) GetJob(ctx context.Context, req *desc.GetJobRequest) (*desc.GetJobResponse, error) {
	jobInfo, err := s.module.GetJob(ctx, &domain.GetJobRequest{
		JobID: req.JobID,
	})
	if err != nil {
		return nil, err
	}
	return &desc.GetJobResponse{
		JobInfo: &desc.JobInfo{
			JobID:        jobInfo.JobID,
			Title:        jobInfo.Title,
			Salary:       jobInfo.Salary,
			Category:     jobInfo.Category,
			Descr:        jobInfo.Descr,
			Location:     jobInfo.Location,
			ContactNum:   jobInfo.ContactNum,
			ContactEmail: jobInfo.ContactEmail,
			Expiry:       ConvertToProtoTimestamp(jobInfo.Expiry),
			Skills:       jobInfo.Skills,
			Status:       jobInfo.Status,
			Date:         ConvertToProtoTimestamp(jobInfo.DatePosted),
			Count:        int32(jobInfo.ApplicantsCount),
		},
	}, nil
}

func (s *service) AddJob(ctx context.Context, req *desc.AddJobRequest) (*desc.AddJobResponse, error) {
	jobID, err := s.module.AddJob(ctx, &domain.AddJobRequest{
		Title:       req.Title,
		Salary:      req.Salary,
		Category:    req.Category,
		Descr:       req.Descr,
		Location:    req.Location,
		ContactNum:  req.ContactNum,
		ContactEmai: req.ContactEmai,
		Expiry:      ConvertTimestamp(req.Expiry),
		Skills:      req.Skills,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &desc.AddJobResponse{
		JobID: jobID,
	}, nil
}

func (s *service) DeleteJob(ctx context.Context, req *desc.DeleteJobRequest) (*emptypb.Empty, error) {
	err := s.module.DeleteJob(ctx, &domain.DeleteJobRequest{
		JobID: req.JobID,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *service) UpdateJob(ctx context.Context, req *desc.UpdateJobRequest) (*emptypb.Empty, error) {
	err := s.module.UpdateJob(ctx, &domain.UpdateJobRequest{
		JobID: req.JobID,
		JobInfo: &domain.JobInfo{
			Title:           req.JobInfo.Title,
			Salary:          req.JobInfo.Salary,
			Category:        req.JobInfo.Category,
			Descr:           req.JobInfo.Descr,
			Location:        req.JobInfo.Location,
			ContactNum:      req.JobInfo.ContactNum,
			ContactEmail:    req.JobInfo.ContactEmail,
			Expiry:          ConvertTimestamp(req.JobInfo.Expiry),
			Skills:          req.JobInfo.Skills,
			Status:          req.JobInfo.Status,
			DatePosted:      ConvertTimestamp(req.JobInfo.Date),
			ApplicantsCount: int32(req.JobInfo.Count),
		},
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *service) ListJobs(ctx context.Context, _ *emptypb.Empty) (*desc.ListJobResponse, error) {
	jobInfos, err := s.module.ListJobs(ctx)
	if err != nil {
		return nil, err
	}
	var jobs []*desc.JobInfo
	for _, jobInfo := range jobInfos {
		jobs = append(jobs, &desc.JobInfo{
			JobID:        jobInfo.JobID,
			Title:        jobInfo.Title,
			Salary:       jobInfo.Salary,
			Category:     jobInfo.Category,
			Descr:        jobInfo.Descr,
			Location:     jobInfo.Location,
			ContactNum:   jobInfo.ContactNum,
			ContactEmail: jobInfo.ContactEmail,
			Expiry:       ConvertToProtoTimestamp(jobInfo.Expiry),
			Skills:       jobInfo.Skills,
			Status:       jobInfo.Status,
			Date:         ConvertToProtoTimestamp(jobInfo.DatePosted),
			Count:        int32(jobInfo.ApplicantsCount),
		})
	}
	return &desc.ListJobResponse{
		Jobs: jobs,
	}, nil
}

func (s *service) Apply(ctx context.Context, req *desc.ApplyRequest) (*emptypb.Empty, error) {
	err := s.module.ApplyJob(ctx, &domain.ApplyRequest{
		UserID: req.UserID,
		JobID:  req.JobID,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *service) ApplyHistory(ctx context.Context, req *desc.ApplyHistoryRequest) (*desc.ApplyHistoryResponse, error) {
	jobInfos, err := s.module.GetAppliedJobs(ctx, &domain.ApplyHistoryRequest{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	var jobs []*desc.JobInfo
	for _, jobInfo := range jobInfos {
		jobs = append(jobs, &desc.JobInfo{
			JobID:        jobInfo.JobID,
			Title:        jobInfo.Title,
			Salary:       jobInfo.Salary,
			Category:     jobInfo.Category,
			Descr:        jobInfo.Descr,
			Location:     jobInfo.Location,
			ContactNum:   jobInfo.ContactNum,
			ContactEmail: jobInfo.ContactEmail,
			Expiry:       ConvertToProtoTimestamp(jobInfo.Expiry),
			Skills:       jobInfo.Skills,
			Status:       jobInfo.Status,
			Date:         ConvertToProtoTimestamp(jobInfo.DatePosted),
			Count:        int32(jobInfo.ApplicantsCount),
		})
	}
	return &desc.ApplyHistoryResponse{
		Jobs: jobs,
	}, nil
}

func ConvertTimestamp(timestamp *timestamp.Timestamp) *time.Time {
	if timestamp == nil {
		return nil
	}
	t := time.Unix(timestamp.Seconds, int64(timestamp.Nanos)).UTC()
	return &t
}
func ConvertToProtoTimestamp(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	ts, _ := ptypes.TimestampProto(*t)
	return ts
}
