package routes

import (
	"movie-watchlist-backend/internal/cache"
	"movie-watchlist-backend/internal/config"
	"movie-watchlist-backend/internal/external"
	"movie-watchlist-backend/internal/handlers"
	"movie-watchlist-backend/internal/middleware"
	"movie-watchlist-backend/internal/repositories"
	"movie-watchlist-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	// Initialize Dependencies
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	tmdbClient := external.NewTMDBClient(cfg)
	movieRepo := repositories.NewMovieRepository()
	movieService := services.NewMovieService(tmdbClient, cache.Client, cfg, movieRepo)
	movieHandler := handlers.NewMovieHandler(movieService)

	watchlistRepo := repositories.NewWatchlistRepository()
	watchlistService := services.NewWatchlistService(watchlistRepo)
	watchlistHandler := handlers.NewWatchlistHandler(watchlistService)

	ratingRepo := repositories.NewRatingRepository()
	recService := services.NewRecommendationService(ratingRepo, tmdbClient)
	ratingHandler := handlers.NewRatingHandler(recService)

	newsService := services.NewNewsService()
	newsHandler := handlers.NewNewsHandler(newsService)

	api := r.Group("/api/v1")
	{
		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Movie Routes (Public)
		movies := api.Group("/movies")
		{
			movies.GET("/search", movieHandler.Search)
			movies.GET("/:id", movieHandler.GetDetails)
		}

		// News Routes (Public)
		api.GET("/news", newsHandler.GetNews)

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.RequireAuth(cfg))
		{
			protected.GET("/profile", func(c *gin.Context) {
				userID := c.MustGet("userID").(uint)
				user, err := userRepo.FindByID(userID)
				if err != nil {
					c.JSON(404, gin.H{"error": "User not found"})
					return
				}
				c.JSON(200, user)
			})

			// Watchlist Routes
			watchlist := protected.Group("/watchlist")
			{
				watchlist.GET("", watchlistHandler.GetUserList)
				watchlist.POST("", watchlistHandler.Add)
				watchlist.PUT("/:id", watchlistHandler.UpdateStatus)
				watchlist.DELETE("/:id", watchlistHandler.Remove)
			}

			// Ratings Routes
			ratings := protected.Group("/ratings")
			{
				ratings.POST("", ratingHandler.AddRating)
				ratings.GET("/recommendations", ratingHandler.GetRecommendations)
			}
		}
	}
}
