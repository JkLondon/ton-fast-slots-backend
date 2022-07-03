package repository

import (
	"Casino/config"
	"Casino/internal/models"
	"Casino/internal/user"
	"Casino/pkg/logger"
	"Casino/pkg/utils"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
)

type userRepo struct {
	db     *sqlx.DB
	config *config.Config
	logger logger.Logger
}

func NewUserRepository(db *sqlx.DB, config *config.Config, logger logger.Logger) user.Repository {
	return &userRepo{db: db, config: config, logger: logger}
}

func (u *userRepo) RegisterUser(ctx context.Context, params models.RegisterUserArgs) (err error) {
	ctx, span := otel.Tracer("").Start(ctx, "userRepo.RegisterUser")
	defer span.End()

	_, err = u.db.ExecContext(ctx, queryCreateUser, params.TGID, params.WalletAddress, params.WalletSeed)
	if err != nil {
		return errors.Wrapf(err, "userRepo.RegisterUser.ExecContext(params: %s)", utils.GetStructJSON(params))
	}

	return nil
}

func (u *userRepo) GetSelfInfo(ctx context.Context, tgID int64) (result models.GetSelfInfoResponse, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "userRepo.GetSelfInfo")
	defer span.End()

	err = u.db.GetContext(ctx, &result, queryGetSelfInfo, tgID)
	if err != nil {
		return result, errors.Wrapf(err, "userRepo.GetSelfInfo.GetContext(tgID: %d)", tgID)
	}

	return result, nil
}
