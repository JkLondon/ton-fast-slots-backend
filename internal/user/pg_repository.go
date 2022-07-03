package user

import (
	"Casino/internal/models"
	"context"
)

type Repository interface {
	RegisterUser(ctx context.Context, params models.RegisterUserArgs) (err error)
	GetSelfInfo(ctx context.Context, tgID int64) (result models.GetSelfInfoResponse, err error)
}
