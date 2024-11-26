package cmd

import (
	"github.com/rs/zerolog"

	"github.com/edjubert/tiny-url/internal/pkg/config"
	"github.com/edjubert/tiny-url/internal/pkg/database"
	"github.com/edjubert/tiny-url/pkg/logger"
)

type Handlers struct {
	logger *zerolog.Logger
	config *config.Configuration
	db     *database.Database
}

func New(config *config.Configuration, db *database.Database) *Handlers {
	return &Handlers{config: config, db: db, logger: logger.New(config.Server.LogLevel)}
}
