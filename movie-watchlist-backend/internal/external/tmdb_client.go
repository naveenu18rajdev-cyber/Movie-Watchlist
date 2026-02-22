package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"movie-watchlist-backend/internal/config"
)

type TMDBClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewTMDBClient(cfg *config.Config) *TMDBClient {
	return &TMDBClient{
		APIKey:  cfg.TMDBApiKey,
		BaseURL: "https://api.themoviedb.org/3",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SearchMovies perform a TMDB search
func (c *TMDBClient) SearchMovies(query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/search/movie?api_key=%s&query=%s", c.BaseURL, c.APIKey, query)
	return c.doRequest(url)
}

// GetMovieDetails fetch rich details for a specific movie
func (c *TMDBClient) GetMovieDetails(movieID int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/movie/%d?api_key=%s&append_to_response=release_dates,watch/providers", c.BaseURL, movieID, c.APIKey)
	return c.doRequest(url)
}

func (c *TMDBClient) doRequest(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("TMDB API returned status: %d", res.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
