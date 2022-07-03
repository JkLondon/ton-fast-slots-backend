package config

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	TON           TON
	Server        Server
	Middleware    Middleware
	Postgres      Postgres
	Session       Session
	TGBot         TGBot
	Logger        Logger
	OpenTelemetry OpenTelemetry
}

type TON struct {
	BaseURL string
}

type Server struct {
	Host                        string `validate:"required"`
	Port                        string `validate:"required"`
	IPHeader                    string `validate:"required"`
	AllowOrigins                string `validate:"required,min=1"`
	AllowHeaders                string `validate:"required,min=1"`
	AllowMethods                string `validate:"required,min=1"`
	ShowUnknownErrorsInResponse bool
	Production                  bool
	RequestTimeout              time.Duration `validate:"required"`
}

type Middleware struct {
	Enable bool
	TGID   int64
}

type Postgres struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SSLMode  string `validate:"required"`
	PgDriver string `validate:"required"`
	Settings struct {
		MaxOpenConns    int           `validate:"required,min=1"`
		ConnMaxLifetime time.Duration `validate:"required,min=1"`
		MaxIdleConns    int           `validate:"required,min=1"`
		ConnMaxIdleTime time.Duration `validate:"required,min=1"`
	}
}

type Session struct {
	RedisLifetime int           `validate:"required"` // Day
	CacheLifetime time.Duration `validate:"required"` // Hour
	CacheCleanUp  time.Duration `validate:"required"` // Minute
}

type TGBot struct {
	Token string `validate:"required"`
}

type Logger struct {
	Level          string `validate:"required"`
	SkipFrameCount int
	InFile         bool
	FilePath       string
	InTG           bool
	ChatID         int64
	TGToken        string
	AlertUsers     []string
}

type OpenTelemetry struct {
	URL         string `validate:"required"`
	ServiceName string `validate:"required"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	err = validator.New().Struct(c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
