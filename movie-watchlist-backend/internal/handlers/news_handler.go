package handlers

import (
	"net/http"

	"movie-watchlist-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	newsService *services.NewsService
}

func NewNewsHandler(service *services.NewsService) *NewsHandler {
	return &NewsHandler{
		newsService: service,
	}
}

func (h *NewsHandler) GetNews(c *gin.Context) {
	news, err := h.newsService.GetLatestNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}

	c.JSON(http.StatusOK, news)
}
