package handlers

import (
	"net/http"
	"strconv"

	"movie-watchlist-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type WatchlistHandler struct {
	watchlistService *services.WatchlistService
}

func NewWatchlistHandler(service *services.WatchlistService) *WatchlistHandler {
	return &WatchlistHandler{
		watchlistService: service,
	}
}

type AddWatchlistReq struct {
	MovieID uint   `json:"movie_id" binding:"required"`
	Status  string `json:"status"`
}

func (h *WatchlistHandler) Add(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var req AddWatchlistReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.watchlistService.AddToWatchlist(userID, req.MovieID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to watchlist"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *WatchlistHandler) GetUserList(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	list, err := h.watchlistService.GetWatchlist(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve watchlist"})
		return
	}

	c.JSON(http.StatusOK, list)
}

type UpdateStatusReq struct {
	Status string `json:"status" binding:"required"`
}

func (h *WatchlistHandler) UpdateStatus(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	itemIDStr := c.Param("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var req UpdateStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.watchlistService.UpdateStatus(uint(itemID), userID, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}

func (h *WatchlistHandler) Remove(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	itemIDStr := c.Param("id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	if err := h.watchlistService.RemoveFromWatchlist(uint(itemID), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from watchlist"})
}
