package api

import (
	"github.com/edjubert/tiny-url/internal/api/cmd"
	"github.com/edjubert/tiny-url/internal/pkg/config"
	"github.com/edjubert/tiny-url/internal/pkg/database"
	"github.com/edjubert/tiny-url/pkg/logger"
)

func setup(configPath string) (*Api, error) {
	conf, err := config.New(configPath)
	if err != nil {
		return nil, err
	}

	log := logger.New(conf.Server.LogLevel)
	log.Info().Msg(conf.Server.LogLevel)

	db, err := database.New(conf)
	if err != nil {
		return nil, err
	}

	server := cmd.Router(conf, db)

	return &Api{
		Config:   conf,
		Database: db,
		Server:   server,
		Logger:   log,
	}, nil
}
