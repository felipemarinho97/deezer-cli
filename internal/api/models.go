package api

import (
	"fmt"
	"time"
)

type Track struct {
	ID                   int64   `json:"id"`
	Title                string  `json:"title"`
	TitleShort          string  `json:"title_short"`
	TitleVersion        string  `json:"title_version"`
	Link                string  `json:"link"`
	Duration            int     `json:"duration"`
	Rank                int     `json:"rank"`
	ExplicitLyrics      bool    `json:"explicit_lyrics"`
	Preview             string  `json:"preview"`
	BPM                 float64 `json:"bpm"`
	Gain                float64 `json:"gain"`
	Artist              Artist  `json:"artist"`
	Album               Album   `json:"album"`
	Type                string  `json:"type"`
}

type Album struct {
	ID                   int64   `json:"id"`
	Title                string  `json:"title"`
	Link                string  `json:"link"`
	Cover               string  `json:"cover"`
	CoverSmall          string  `json:"cover_small"`
	CoverMedium         string  `json:"cover_medium"`
	CoverBig            string  `json:"cover_big"`
	CoverXL             string  `json:"cover_xl"`
	GenreID             int     `json:"genre_id"`
	NbTracks            int     `json:"nb_tracks"`
	ReleaseDate         string  `json:"release_date"`
	RecordType          string  `json:"record_type"`
	Tracklist           string  `json:"tracklist"`
	ExplicitLyrics      bool    `json:"explicit_lyrics"`
	Artist              Artist  `json:"artist"`
	Type                string  `json:"type"`
}

type Artist struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Link             string `json:"link"`
	Picture          string `json:"picture"`
	PictureSmall     string `json:"picture_small"`
	PictureMedium    string `json:"picture_medium"`
	PictureBig       string `json:"picture_big"`
	PictureXL        string `json:"picture_xl"`
	NbAlbum          int    `json:"nb_album"`
	NbFan            int    `json:"nb_fan"`
	Radio            bool   `json:"radio"`
	Tracklist        string `json:"tracklist"`
	Type             string `json:"type"`
}

type Playlist struct {
	ID               int64         `json:"id"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	Duration         int           `json:"duration"`
	Public           bool          `json:"public"`
	IsLovedTrack     bool          `json:"is_loved_track"`
	Collaborative    bool          `json:"collaborative"`
	NbTracks         int           `json:"nb_tracks"`
	Fans             int           `json:"fans"`
	Link             string        `json:"link"`
	Picture          string        `json:"picture"`
	PictureSmall     string        `json:"picture_small"`
	PictureMedium    string        `json:"picture_medium"`
	PictureBig       string        `json:"picture_big"`
	PictureXL        string        `json:"picture_xl"`
	Checksum         string        `json:"checksum"`
	Creator          *User         `json:"creator"`
	Tracks           *TracksData   `json:"tracks"`
	Type             string        `json:"type"`
	CreationDate     string        `json:"creation_date"`
}

type User struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Tracklist        string `json:"tracklist"`
	Type             string `json:"type"`
}

type TracksData struct {
	Data []Track `json:"data"`
}

type TrackSearchResult struct {
	Data  []Track `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type AlbumSearchResult struct {
	Data  []Album `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type ArtistSearchResult struct {
	Data  []Artist `json:"data"`
	Total int      `json:"total"`
	Next  string   `json:"next"`
}

type PlaylistSearchResult struct {
	Data  []Playlist `json:"data"`
	Total int        `json:"total"`
	Next  string     `json:"next"`
}

type TracksResult struct {
	Data  []Track `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

type AlbumsResult struct {
	Data  []Album `json:"data"`
	Total int     `json:"total"`
	Next  string  `json:"next"`
}

func (t Track) GetID() int64 {
	return t.ID
}

func (t Track) GetTitle() string {
	return t.Title
}

func (t Track) GetArtistName() string {
	return t.Artist.Name
}

func (t Track) GetAlbumTitle() string {
	return t.Album.Title
}

func (t Track) GetDurationFormatted() string {
	duration := time.Duration(t.Duration) * time.Second
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func (a Album) GetID() int64 {
	return a.ID
}

func (a Album) GetTitle() string {
	return a.Title
}

func (a Album) GetArtistName() string {
	return a.Artist.Name
}

func (a Artist) GetID() int64 {
	return a.ID
}

func (a Artist) GetName() string {
	return a.Name
}

func (p Playlist) GetID() int64 {
	return p.ID
}

func (p Playlist) GetTitle() string {
	return p.Title
}

func (p Playlist) GetCreatorName() string {
	if p.Creator != nil {
		return p.Creator.Name
	}
	return ""
}