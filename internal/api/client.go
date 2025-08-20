package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/felipemarinho97/deezer-cli/internal/cache"
)

const (
	BaseURL = "https://api.deezer.com"
)

type Client struct {
	httpClient  *http.Client
	baseURL     string
	rateLimiter *time.Ticker
	cache       *cache.Cache
}

func NewClient() *Client {
	return &Client{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		baseURL:     BaseURL,
		rateLimiter: time.NewTicker(50 * time.Millisecond), // 20 requests per second max
		cache:       cache.New(5*time.Minute, 10*time.Minute),
	}
}

func NewClientWithCache(cacheEnabled bool, ttl time.Duration) *Client {
	client := &Client{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		baseURL:     BaseURL,
		rateLimiter: time.NewTicker(50 * time.Millisecond),
	}
	
	if cacheEnabled {
		client.cache = cache.New(ttl, ttl*2)
	}
	
	return client
}

func (c *Client) get(endpoint string, params url.Values) ([]byte, error) {
	cacheKey := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	
	if c.cache != nil {
		var cachedData []byte
		if c.cache.Get(cacheKey, &cachedData) {
			return cachedData, nil
		}
	}
	
	<-c.rateLimiter.C

	fullURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	if params != nil && len(params) > 0 {
		fullURL = fmt.Sprintf("%s?%s", fullURL, params.Encode())
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var errorCheck struct {
		Error *struct {
			Type    string `json:"type"`
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}
	
	if err := json.Unmarshal(body, &errorCheck); err == nil && errorCheck.Error != nil {
		return nil, fmt.Errorf("API error: %s", errorCheck.Error.Message)
	}

	if c.cache != nil {
		c.cache.Set(cacheKey, body)
	}

	return body, nil
}

func (c *Client) SearchTracks(query string, limit int, index int) (*TrackSearchResult, error) {
	params := url.Values{}
	params.Set("q", query)
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if index > 0 {
		params.Set("index", fmt.Sprintf("%d", index))
	}

	data, err := c.get("/search/track", params)
	if err != nil {
		return nil, err
	}

	var result TrackSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) SearchAlbums(query string, limit int, index int) (*AlbumSearchResult, error) {
	params := url.Values{}
	params.Set("q", query)
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if index > 0 {
		params.Set("index", fmt.Sprintf("%d", index))
	}

	data, err := c.get("/search/album", params)
	if err != nil {
		return nil, err
	}

	var result AlbumSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) SearchArtists(query string, limit int, index int) (*ArtistSearchResult, error) {
	params := url.Values{}
	params.Set("q", query)
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if index > 0 {
		params.Set("index", fmt.Sprintf("%d", index))
	}

	data, err := c.get("/search/artist", params)
	if err != nil {
		return nil, err
	}

	var result ArtistSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) SearchPlaylists(query string, limit int, index int) (*PlaylistSearchResult, error) {
	params := url.Values{}
	params.Set("q", query)
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	if index > 0 {
		params.Set("index", fmt.Sprintf("%d", index))
	}

	data, err := c.get("/search/playlist", params)
	if err != nil {
		return nil, err
	}

	var result PlaylistSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) GetTrack(id int64) (*Track, error) {
	endpoint := fmt.Sprintf("/track/%d", id)
	data, err := c.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var track Track
	if err := json.Unmarshal(data, &track); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &track, nil
}

func (c *Client) GetAlbum(id int64) (*Album, error) {
	endpoint := fmt.Sprintf("/album/%d", id)
	data, err := c.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var album Album
	if err := json.Unmarshal(data, &album); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &album, nil
}

func (c *Client) GetArtist(id int64) (*Artist, error) {
	endpoint := fmt.Sprintf("/artist/%d", id)
	data, err := c.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var artist Artist
	if err := json.Unmarshal(data, &artist); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &artist, nil
}

func (c *Client) GetPlaylist(id int64) (*Playlist, error) {
	endpoint := fmt.Sprintf("/playlist/%d", id)
	data, err := c.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var playlist Playlist
	if err := json.Unmarshal(data, &playlist); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &playlist, nil
}

func (c *Client) GetAlbumTracks(id int64, limit int) (*TracksResult, error) {
	endpoint := fmt.Sprintf("/album/%d/tracks", id)
	params := url.Values{}
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}

	data, err := c.get(endpoint, params)
	if err != nil {
		return nil, err
	}

	var result TracksResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) GetArtistAlbums(id int64, limit int) (*AlbumsResult, error) {
	endpoint := fmt.Sprintf("/artist/%d/albums", id)
	params := url.Values{}
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}

	data, err := c.get(endpoint, params)
	if err != nil {
		return nil, err
	}

	var result AlbumsResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *Client) GetArtistTopTracks(id int64, limit int) (*TracksResult, error) {
	endpoint := fmt.Sprintf("/artist/%d/top", id)
	params := url.Values{}
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}

	data, err := c.get(endpoint, params)
	if err != nil {
		return nil, err
	}

	var result TracksResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func FilterByArtist(tracks []Track, artistName string) []Track {
	if artistName == "" {
		return tracks
	}

	var filtered []Track
	lowerArtist := strings.ToLower(artistName)
	
	for _, track := range tracks {
		if strings.ToLower(track.Artist.Name) == lowerArtist {
			filtered = append(filtered, track)
		}
	}
	
	return filtered
}

func FilterAlbumsByArtist(albums []Album, artistName string) []Album {
	if artistName == "" {
		return albums
	}

	var filtered []Album
	lowerArtist := strings.ToLower(artistName)
	
	for _, album := range albums {
		if strings.ToLower(album.Artist.Name) == lowerArtist {
			filtered = append(filtered, album)
		}
	}
	
	return filtered
}