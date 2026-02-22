# Movie Watchlist Backend API

This is the backend REST API for the Movie Watchlist & Recommendation platform, built using Go (Golang) and the Gin web framework.

## Prerequisites
- Go 1.21 or higher
- PostgreSQL
- Redis

## Setup Instructions
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Update the `.env` file with your local Database, Redis, and API Keys.
3. Start the server:
   ```bash
   go run cmd/server/main.go
   ```
