package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/edjubert/tiny-url/internal/pkg/database"
)

type CreateForm struct {
	Url string `json:"url" binding:"required"`
}

func (h *Handlers) Create(c *gin.Context) {
	var form CreateForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := url.Parse(form.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if u == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is nil"})
		return
	}

	if u.Scheme == "" || u.Host == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not a valid url"})
		return
	}

	slug, expirationDate, err := database.Create(c.Request.Context(), h.db, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"url":             fmt.Sprintf("%s://%s:%s/%s", h.config.Server.Scheme, h.config.Server.Host, h.config.Server.Port, slug),
			"expiration_date": expirationDate,
		},
	)
}
