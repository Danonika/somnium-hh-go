package user

import (
	"context"
	"somnium/internal/domain"
)

type UserModule struct {
	db domain.SomniumSystemPostgre
}

func NewUserModule(db domain.SomniumSystemPostgre) *UserModule {
	return &UserModule{
		db: db,
	}
}

func (am *UserModule) UpdateUser(ctx context.Context, input *domain.UpdateUserRequest) error {
	return am.db.UpdateUser(ctx, input)
}

func (am *UserModule) GetUser(ctx context.Context, input *domain.GetUserRequest) (*domain.UserInfo, error) {
	return am.db.GetUser(ctx, input)
}
