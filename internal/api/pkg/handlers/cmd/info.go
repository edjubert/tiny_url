package cmd

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/edjubert/tiny-url/internal/pkg/database"
)

type GetClickedForm struct {
	Slug string `form:"slug" binding:"required"`
}

func (h *Handlers) Info(c *gin.Context) {
	var form GetClickedForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := database.GetUrlBySlug(c.Request.Context(), h.db, form.Slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"id":      url.Id,
			"slug":    url.Slug,
			"url":     url.Url,
			"clicked": url.Clicked,

			"created_at": url.CreatedAt,
			"expires_at": url.ExpiresAt,
		},
	)
}
