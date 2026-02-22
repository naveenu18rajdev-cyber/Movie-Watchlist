# Movie Watchlist Platform

A full-stack web application designed for discovering movies, managing personal watchlists, and receiving recommendations.

## Project Structure

The project is divided into two main components:
- `/movie-watchlist-backend`: A REST API built with Go (Golang) and PostgreSQL.
- `/movie-watchlist-frontend`: A modern Single Page Application (SPA) built with React and Vite.

## Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) installed for the backend
- [Node.js](https://nodejs.org/) installed for the frontend
- [PostgreSQL](https://www.postgresql.org/) database (or Docker to run via docker-compose)

### Running the Project

You can use Docker Compose to easily spin up the database, or run each part individually. Refer to the specific `README.md` inside each subfolder for detailed, step-by-step setup instructions:

* **[Backend Details & Setup](./movie-watchlist-backend/README.md)**
* **[Frontend Details & Setup](./movie-watchlist-frontend/README.md)**

## Features
- User authentication (JWT based)
- Search and discover popular movies
- Add movies to custom watchlists
- Rate and review movies
- Get personalized movie recommendations

## External APIs Used
- **[TMDB (The Movie Database) API](https://developer.themoviedb.org/docs/getting-started)**: Used extensively in the backend for fetching detailed movie information, searching titles, getting trending lists, and discovering new films. A valid API Key is required in the backend `.env` file to enable this functionality.
