package domain

import (
	"context"
	"fmt"
	"strings"
)

type ctxKey int

const (
	ctxToken ctxKey = iota
	ctxClaims
)

var (
	ErrorEmptyToken       = fmt.Errorf("unauthorized token")
	ErrorEmptyClaims      = fmt.Errorf("unauthorized claims")
	ErrorPermissionDenied = fmt.Errorf("permission denied")
)

func CleanEmail(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	return email
}

func WithClaims[T any](ctx context.Context, claims *T) context.Context {
	return context.WithValue(ctx, ctxClaims, claims)
}

func ExtractClaims[T any](ctx context.Context) (*T, error) {
	claims, ok := ctx.Value(ctxClaims).(*T)
	if !ok {
		return nil, ErrorEmptyClaims
	}
	return claims, nil
}
