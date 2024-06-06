package auth

import (
	"context"
	"somnium/internal/domain"
)

type AuthenticationModule struct {
	db domain.SomniumSystemPostgre
}

func NewAuthenticationModule(db domain.SomniumSystemPostgre) *AuthenticationModule {
	return &AuthenticationModule{
		db: db,
	}
}

func (am *AuthenticationModule) SignIn(ctx context.Context, input *domain.SignInRequest) (*domain.UserClaims, error) {
	userID, err := am.db.SignIn(ctx, input)
	if err != nil {
		return nil, err
	}
	roles, err := am.GetRoles(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &domain.UserClaims{
		UserID: userID,
		Roles:  roles,
	}, nil
}
func (am *AuthenticationModule) SignUp(ctx context.Context, input *domain.SignUpRequest) (*domain.UserClaims, error) {
	userID, err := am.db.SignUp(ctx, input)
	if err != nil {
		return nil, err
	}
	roles, err := am.GetRoles(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &domain.UserClaims{
		UserID: userID,
		Roles:  roles,
	}, nil
}

func (am *AuthenticationModule) GetRoles(ctx context.Context, userID string) ([]string, error) {
	return am.db.GetRoles(ctx, userID)
}
