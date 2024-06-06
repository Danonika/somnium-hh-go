package somniumsystem

import (
	"context"
	"somnium/internal/domain"

	desc "somnium/pkg/api/somnium/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	err := s.module.UpdateUser(ctx, &domain.UpdateUserRequest{
		UserID: req.UserID,
		UserInfo: domain.UserInfo{
			Email:       req.UserInfo.Email,
			Name:        req.UserInfo.Name,
			ContactNum:  req.UserInfo.ContactNum,
			ResumeTitle: req.UserInfo.ResumeTitle,
			ResumeLink:  req.UserInfo.ResumeLink,
			Skills:      req.UserInfo.Skills,
		},
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *service) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	userInfo, err := s.module.GetUser(ctx, &domain.GetUserRequest{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &desc.GetUserResponse{
		UserInfo: &desc.UserInfo{
			Email:       userInfo.Email,
			Name:        userInfo.Name,
			ContactNum:  userInfo.ContactNum,
			ResumeTitle: userInfo.ResumeTitle,
			ResumeLink:  userInfo.ResumeLink,
			Skills:      userInfo.Skills,
		},
	}, nil
}
