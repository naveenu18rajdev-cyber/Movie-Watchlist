package handlers

import (
	"net/http"
	"strconv"

	"movie-watchlist-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieService *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
	return &MovieHandler{
		movieService: service,
	}
}

func (h *MovieHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required for searching"})
		return
	}

	results, err := h.movieService.SearchMovies(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search movies"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *MovieHandler) GetDetails(c *gin.Context) {
	idStr := c.Param("id")
	movieID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	details, err := h.movieService.GetMovieDetails(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie details"})
		return
	}

	c.JSON(http.StatusOK, details)
}
