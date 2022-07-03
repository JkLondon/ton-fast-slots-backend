package usecase

import (
	"Casino/config"
	"Casino/internal/game"
	"Casino/internal/models"
	"Casino/pkg/logger"
	"Casino/pkg/slot"
	"Casino/pkg/tonapi"
	"context"
	"strconv"

	"github.com/patrickmn/go-cache"
	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"
)

type gameUC struct {
	cfg         *config.Config
	bot         *telebot.Bot
	tonSDK      tonapi.SDK
	slotManager slot.Manager
	gameSession *cache.Cache
	logger      logger.Logger
}

func NewGameUseCase(
	cfg *config.Config,
	tonSDK tonapi.SDK,
	log logger.Logger,
) game.UseCase {
	return &gameUC{
		cfg:         cfg,
		slotManager: slot.InitManager(),
		gameSession: cache.New(cache.NoExpiration, cache.NoExpiration),
		tonSDK:      tonSDK,
		logger:      log,
	}
}

func (g *gameUC) StartGame(ctx context.Context, tgID int64) (result models.StartGameResult, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "gameUC.StartGame")
	defer span.End()

	_, exists := g.gameSession.Get(strconv.FormatInt(tgID, 10))
	if exists {
		return models.StartGameResult{
			Success: false,
			Cause:   "already exists session",
		}, nil
	}

	g.gameSession.Set(strconv.FormatInt(tgID, 10), struct{}{}, cache.NoExpiration)

	return models.StartGameResult{
		Success: true,
	}, nil
}

func (g *gameUC) RoundSlot(
	ctx context.Context,
	tgID int64,
	params models.RoundSlotParams,
) (result models.RoundSlotResult, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "gameUC.RoundSlot")
	defer span.End()
	_, exists := g.gameSession.Get(strconv.FormatInt(tgID, 10))
	if !exists {
		return models.RoundSlotResult{
			Success: false,
			Cause:   "game does not exists session",
		}, nil
	}

	multiply, scrollRes := g.slotManager.Scroll()
	if multiply == 0 {
		return models.RoundSlotResult{
			Success:      true,
			WinAmount:    -params.Amount,
			ScrollResult: scrollRes,
		}, nil
	}

	return models.RoundSlotResult{
		Success:      true,
		WinAmount:    params.Amount * multiply,
		ScrollResult: scrollRes,
	}, nil
}

func (g *gameUC) EndGame(ctx context.Context, tgID int64) (result models.EndGameResult, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "gameUC.EndGame")
	defer span.End()

	_, exists := g.gameSession.Get(strconv.FormatInt(tgID, 10))
	if !exists {
		return models.EndGameResult{
			Success: false,
			Cause:   "game does not exists session",
		}, nil
	}
	return result, nil
}
