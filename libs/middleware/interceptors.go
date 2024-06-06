package mw

import (
	"context"
	"somnium/internal/domain"
	"somnium/libs/jwt"
	"strings"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (mw *Middleware) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.FullMethod == "/somniumsystem.v1.SomniumService/SignUp" || info.FullMethod == "/somniumsystem.v1.SomniumService/SignIn" {
		return handler(ctx, req)
	}
	print(info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	token := extractToken(md, "authorization")
	if token == "" {
		return nil, status.Errorf(codes.Unauthenticated, "missing token")
	}
	claims, err := jwt.Parse[domain.UserClaims](mw.jc, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}
	if !hasPermission(claims.Roles, info.FullMethod) {
		return nil, status.Error(codes.Unauthenticated, "access denied")
	}
	newCtx := domain.WithClaims(ctx, claims)
	return handler(newCtx, req)
}
func (mw *Middleware) ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err == pgx.ErrNoRows {
		return resp, status.Errorf(codes.NotFound, "The requested resource was not found")
	}
	return resp, err
}
func extractToken(md metadata.MD, key string) string {
	vals := md[key]
	if len(vals) == 0 {
		return ""
	}
	return strings.TrimPrefix(vals[0], "Bearer ")
}

func hasPermission(roles []string, method string) bool {
	if containsString(roles, "admin") {
		return true
	}
	if containsString(roles, "user") {
		if method == "/somniumsystem.v1.SomniumService/Apply" || method == "/somniumsystem.v1.SomniumService/ApplyHistory" || method == "/somniumsystem.v1.SomniumService/GetJob" || method == "/somniumsystem.v1.SomniumService/AddSkill" || method == "/somniumsystem.v1.SomniumService/SkillPool" || method == "/somniumsystem.v1.SomniumService/GetUser" || method == "/somniumsystem.v1.SomniumService/UpdateUser" {
			return true
		}
	}
	return false
}

func containsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
