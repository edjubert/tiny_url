package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/edjubert/tiny-url/internal/pkg/database"
)

func (h *Handlers) SlugRedirect(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slug required"})
		return
	}

	url, err := database.GetUrlBySlug(c.Request.Context(), h.db, slug)
	if err != nil {
		c.JSON(
			404, gin.H{
				"error":   "not found",
				"message": "You can create your own tiny url with a POST request to http://localhost:3000/create with the field 'url'",
			},
		)
		return
	}

	if err := database.AddClicked(c.Request.Context(), h.db, slug); err != nil {
		c.JSON(500, gin.H{"error": "something went wrong"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url.Url)
}
