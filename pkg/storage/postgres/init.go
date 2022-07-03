package postgres

import (
	"Casino/config"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

func InitPsqlDB(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.DBName,
		c.Postgres.SSLMode)

	database, err := sqlx.Connect(c.Postgres.PgDriver, connectionUrl)
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(c.Postgres.Settings.MaxOpenConns)
	database.SetConnMaxLifetime(c.Postgres.Settings.ConnMaxLifetime * time.Second)
	database.SetMaxIdleConns(c.Postgres.Settings.MaxIdleConns)
	database.SetConnMaxIdleTime(c.Postgres.Settings.ConnMaxIdleTime * time.Second)

	if err = database.Ping(); err != nil {
		return nil, err
	}
	return database, nil
}
