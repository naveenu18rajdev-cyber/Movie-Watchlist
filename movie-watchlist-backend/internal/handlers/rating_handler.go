package handlers

import (
	"net/http"

	"movie-watchlist-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	recService *services.RecommendationService
}

func NewRatingHandler(service *services.RecommendationService) *RatingHandler {
	return &RatingHandler{
		recService: service,
	}
}

type AddRatingReq struct {
	MovieID uint   `json:"movie_id" binding:"required"`
	Score   int    `json:"score" binding:"required,min=1,max=5"`
	Review  string `json:"review"`
}

func (h *RatingHandler) AddRating(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var req AddRatingReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating, err := h.recService.AddRating(userID, req.MovieID, req.Score, req.Review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add rating"})
		return
	}

	c.JSON(http.StatusCreated, rating)
}

func (h *RatingHandler) GetRecommendations(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	recs, err := h.recService.GenerateRecommendations(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate recommendations"})
		return
	}

	c.JSON(http.StatusOK, recs)
}
