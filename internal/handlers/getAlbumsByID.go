package handlers

import (
	"net/http"

	"github.com/aventhis/go-vinilStore-api/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, idAlbums := range models.Albums {
		if idAlbums.ID == id {
			c.IndentedJSON(http.StatusOK, idAlbums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"info": "album not found"})
}
