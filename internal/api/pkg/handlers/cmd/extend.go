package cmd

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/edjubert/tiny-url/internal/pkg/database"
)

type ExtendForm struct {
	Slug string `json:"slug" binding:"required"`
}

func (h *Handlers) Extend(c *gin.Context) {
	var form ExtendForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newExpirationDate, err := database.ExtendSlug(c.Request.Context(), h.db, form.Slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{"expires_at": newExpirationDate.Format(time.RFC3339)},
	)
}
