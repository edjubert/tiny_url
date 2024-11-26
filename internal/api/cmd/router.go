package cmd

import (
	"github.com/gin-gonic/gin"

	"github.com/edjubert/tiny-url/internal/api/pkg/handlers/cmd"
	"github.com/edjubert/tiny-url/internal/pkg/config"
	"github.com/edjubert/tiny-url/internal/pkg/database"
)

func Router(conf *config.Configuration, db *database.Database) *gin.Engine {
	switch conf.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	app := gin.New()

	app.Use(gin.Recovery())

	h := cmd.New(conf, db)

	app.POST("/create", h.Create)
	return app
}
