package somniumsystem

import (
	"context"
	"somnium/internal/domain"
	desc "somnium/pkg/api/somnium/v1"
)

func (s *service) SignIn(ctx context.Context, req *desc.SignInRequest) (*desc.SignInResponse, error) {
	response, err := s.module.SignIn(ctx, &domain.SignInRequest{
		Email:    domain.CleanEmail(req.Email),
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	creds, err := s.module.Jwtcli.Create(ctx, response)
	if err != nil {
		return nil, err
	}

	return &desc.SignInResponse{
		AccessToken: creds.AccessToken,
	}, nil
}

func (s *service) SignUp(ctx context.Context, req *desc.SignUpRequest) (*desc.SignUpResponse, error) {
	response, err := s.module.SignUp(ctx, &domain.SignUpRequest{
		Email:    domain.CleanEmail(req.Email),
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	creds, err := s.module.Jwtcli.Create(ctx, response)
	if err != nil {
		return nil, err
	}

	return &desc.SignUpResponse{
		AccessToken: creds.AccessToken,
	}, nil
}
