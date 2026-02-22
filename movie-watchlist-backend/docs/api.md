# Movie Watchlist Platform API Documentation

## Auth Endpoints
### `POST /api/v1/auth/register`
**Body:** `{ "email": "user@example.com", "password": "password123", "name": "John" }`
Registers a new user.

### `POST /api/v1/auth/login`
**Body:** `{ "email": "user@example.com", "password": "password123" }`
Returns a JWT token `{"token": "...", "user": {...}}`.

## Public Movie Endpoints
### `GET /api/v1/movies/search?q={query}`
Searches movies using TMDB. Cached in Redis for 1 hour.

### `GET /api/v1/movies/:id`
Gets full movie details including watch providers (streaming availability). Cached in Redis for 24 hours.

### `GET /api/v1/news`
Fetches latest entertainment news.

## Protected Endpoints (Requires `Authorization: Bearer <token>`)
### `GET /api/v1/profile`
Returns current user profile.

### `GET /api/v1/watchlist`
Returns the user's saved watchlist.

### `POST /api/v1/watchlist`
**Body:** `{ "movie_id": 12345, "status": "Want to Watch" }`
Adds a movie to the watchlist.

### `PUT /api/v1/watchlist/:id`
**Body:** `{ "status": "Watched" }`
Updates the watchlist status.

### `DELETE /api/v1/watchlist/:id`
Removes movie from watchlist.

### `POST /api/v1/ratings`
**Body:** `{ "movie_id": 12345, "score": 5, "review": "Great!" }`
Add a rating for a movie.

### `GET /api/v1/ratings/recommendations`
Generates personalized recommendations based on past ratings.
