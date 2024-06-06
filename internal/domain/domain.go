package domain

import (
	"context"
	"time"
)

const (
	RoleUser   = "user"
	RoleAdmin  = "admin"
	RoleSystem = "system"
)

type SignInRequest struct {
	Email    string
	Password string
}
type SignUpRequest struct {
	Email    string
	Password string
	Attrs    string
}

type UpdateUserRequest struct {
	UserID   string
	UserInfo UserInfo
}

type UserInfo struct {
	Email       string
	Name        string
	ContactNum  string
	ResumeTitle string
	ResumeLink  string
	Skills      []string
}
type GetUserRequest struct {
	UserID string
}
type AddSkillRequest struct {
	Skill string
}

type SkillPoolResponse struct {
	Skills []string
}

type GetJobRequest struct {
	JobID string
}

type GetJobResponse struct {
	JobInfo *JobInfo
}

type AddJobRequest struct {
	Title       string
	Salary      string
	Category    string
	Descr       string
	Location    string
	ContactNum  string
	ContactEmai string
	Expiry      *time.Time
	Skills      []string
	Status      string
}

type AddJobResponse struct {
	JobID string
}

type DeleteJobRequest struct {
	JobID string
}

type UpdateJobRequest struct {
	JobID   string
	JobInfo *JobInfo
}

type ListJobResponse struct {
	Jobs []*JobInfo
}

type JobInfo struct {
	JobID           string
	Title           string
	Salary          string
	Category        string
	Descr           string
	Location        string
	ContactNum      string
	ContactEmail    string
	Expiry          *time.Time
	Skills          []string
	Status          string
	DatePosted      *time.Time
	ApplicantsCount int32
}

type JobSwitcherRequest struct {
	JobID string
}
type ApplyRequest struct {
	UserID string
	JobID  string
}

type ApplyHistoryRequest struct {
	UserID string
}

type ApplyHistoryResponse struct {
	Jobs []JobInfo
}
type UserClaims struct {
	UserID string   `json:"uid"`
	Roles  []string `json:"r,omitempty"`
}

type Credentials struct {
	AccessToken string `json:"access_token"`
}
type CredentialsRepository interface {
	Create(ctx context.Context, claims *UserClaims) (*Credentials, error)
}
type MiddlewareRepository interface {
}

type SomniumSystemPostgre interface {
	//Auth
	SignIn(ctx context.Context, input *SignInRequest) (string, error)
	SignUp(ctx context.Context, input *SignUpRequest) (string, error)
	//User
	GetRoles(ctx context.Context, userID string) ([]string, error)
	UpdateUser(ctx context.Context, input *UpdateUserRequest) error
	GetUser(ctx context.Context, input *GetUserRequest) (*UserInfo, error)
	//Job
	AddSkill(ctx context.Context, skill string) error
	SkillPool(ctx context.Context) ([]string, error)
	AddJob(ctx context.Context, req *AddJobRequest) (string, error)
	GetJob(ctx context.Context, req *GetJobRequest) (*JobInfo, error)
	DeleteJob(ctx context.Context, req *DeleteJobRequest) error
	UpdateJob(ctx context.Context, req *UpdateJobRequest) error
	ListJobs(ctx context.Context) ([]JobInfo, error)
	ApplyJob(ctx context.Context, req *ApplyRequest) error
	GetAppliedJobs(ctx context.Context, req *ApplyHistoryRequest) ([]JobInfo, error)
}
