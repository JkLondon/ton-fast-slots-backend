package user

import (
	"Casino/internal/models"
	"context"
)

type UseCase interface {
	SetBot() error
	GetSelfInfo(ctx context.Context, tgID int64) (result models.GetSelfInfoResponse, err error)
}
