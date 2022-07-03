package usecase

import (
	"Casino/config"
	"Casino/internal/models"
	"Casino/internal/user"
	"Casino/pkg/logger"
	"Casino/pkg/tonapi"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"
)

type userUC struct {
	cfg         *config.Config
	userRepo    user.Repository
	bot         *telebot.Bot
	tonSDK      tonapi.SDK
	userSession *cache.Cache
	logger      logger.Logger
}

func NewUserUseCase(
	cfg *config.Config,
	userRepo user.Repository,
	userSession *cache.Cache,
	tonSDK tonapi.SDK,
	log logger.Logger,
) user.UseCase {
	return &userUC{
		cfg:         cfg,
		userRepo:    userRepo,
		userSession: userSession,
		tonSDK:      tonSDK,
		logger:      log,
	}
}

func (u *userUC) SetBot() error {
	pref := telebot.Settings{
		Token:  u.cfg.TGBot.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return err
	}

	b.Handle("/start", func(c telebot.Context) error {
		ctx, span := otel.Tracer("").Start(context.Background(), "userUC.Bot.start")
		defer span.End()
		err = u.userRepo.RegisterUser(ctx, models.RegisterUserArgs{
			TGID:          c.Chat().ID,
			WalletAddress: uuid.New().String(),
			WalletSeed:    uuid.New().String(),
		})
		if err != nil {
			u.logger.Error(err)
			return c.Send("Internal error")
		}
		return c.Send("Success registered. Use /enter_game command for start the game")
	})

	b.Handle("/enter_game", func(c telebot.Context) error {
		_, span := otel.Tracer("").Start(context.Background(), "userUC.Bot.enter_game")
		defer span.End()

		sessionKey := uuid.New().String()
		u.userSession.Set(sessionKey, c.Chat().ID, cache.NoExpiration)
		return c.Send(fmt.Sprintf("Your key for game is - %s", sessionKey))
	})

	go b.Start()
	return nil
}

func (u *userUC) GetSelfInfo(ctx context.Context, tgID int64) (result models.GetSelfInfoResponse, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "userUC.GetSelfInfo")
	defer span.End()

	result, err = u.userRepo.GetSelfInfo(ctx, tgID)
	if err != nil {
		return result, err
	}

	result.Balance, err = u.tonSDK.GetAddressBalance(result.WalletAddress)
	if err != nil {
		return result, err
	}

	return
}
