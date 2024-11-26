package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/edjubert/tiny-url/internal/pkg/config"
	"github.com/edjubert/tiny-url/internal/pkg/database"
)

type Api struct {
	Logger   *zerolog.Logger
	Config   *config.Configuration
	Database *database.Database
	Server   *gin.Engine
}

func New() (*Api, error) {
	var configPath string

	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	env := os.Getenv("MODE")
	switch env {
	case "dev":
	default:
		configPath = rootPath + "/config/dev.yml"
	}

	a, err := setup(configPath)
	if err != nil {
		return nil, err
	}

	return a, nil
}
