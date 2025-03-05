package handlers

import (
	"net/http"

	"github.com/aventhis/go-vinilStore-api/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}
