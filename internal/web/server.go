package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sebin-pavus/sample/internal/web/handler"
)

func NewServer(router *gin.Engine) {
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.PostAlbums)
}
