package main

import (
	"github.com/aventhis/go-vinilStore-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)

	router.Run("localhost:8080")
}
