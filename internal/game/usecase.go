package game

import (
	"Casino/internal/models"
	"context"
)

type UseCase interface {
	StartGame(ctx context.Context, tgID int64) (result models.StartGameResult, err error)
	RoundSlot(ctx context.Context, tgID int64, params models.RoundSlotParams) (result models.RoundSlotResult, err error)
	EndGame(ctx context.Context, tgID int64) (result models.EndGameResult, err error)
}
