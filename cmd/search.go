package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/felipemarinho97/deezer-cli/internal/api"
	"github.com/felipemarinho97/deezer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	searchType   string
	artistFilter string
	albumFilter  string
	exact        bool
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for tracks, albums, artists, or playlists",
	Long: `Search the Deezer catalog for music content.
	
Examples:
  deezer search "daft punk" --type track
  deezer search "random access memories" --type album --limit 5
  deezer search "get lucky" --artist "daft punk" --exact
  deezer search "chill" --type playlist --output json
  deezer search "madonna" --type artist --ids-only`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		client := api.NewClient()
		formatter := output.NewFormatter(outputFormat, idsOnly, fields)

		switch strings.ToLower(searchType) {
		case "track", "tracks":
			searchTracks(client, query, formatter)
		case "album", "albums":
			searchAlbums(client, query, formatter)
		case "artist", "artists":
			searchArtists(client, query, formatter)
		case "playlist", "playlists":
			searchPlaylists(client, query, formatter)
		default:
			searchAll(client, query, formatter)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchType, "type", "t", "all", "Type to search: track, album, artist, playlist, all")
	searchCmd.Flags().StringVar(&artistFilter, "artist", "", "Filter results by artist name (case-insensitive)")
	searchCmd.Flags().StringVar(&albumFilter, "album", "", "Filter results by album name (case-insensitive)")
	searchCmd.Flags().BoolVar(&exact, "exact", false, "Use exact matching for filters")
}

func searchTracks(client *api.Client, query string, formatter *output.Formatter) {
	result, err := client.SearchTracks(query, limit, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error searching tracks: %v\n", err)
		os.Exit(1)
	}

	tracks := result.Data
	
	if artistFilter != "" {
		tracks = api.FilterByArtist(tracks, artistFilter)
	}
	
	if albumFilter != "" {
		tracks = filterTracksByAlbum(tracks, albumFilter)
	}

	formatter.FormatTracks(tracks)
}

func searchAlbums(client *api.Client, query string, formatter *output.Formatter) {
	result, err := client.SearchAlbums(query, limit, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error searching albums: %v\n", err)
		os.Exit(1)
	}

	albums := result.Data
	
	if artistFilter != "" {
		albums = api.FilterAlbumsByArtist(albums, artistFilter)
	}

	formatter.FormatAlbums(albums)
}

func searchArtists(client *api.Client, query string, formatter *output.Formatter) {
	result, err := client.SearchArtists(query, limit, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error searching artists: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatArtists(result.Data)
}

func searchPlaylists(client *api.Client, query string, formatter *output.Formatter) {
	result, err := client.SearchPlaylists(query, limit, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error searching playlists: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatPlaylists(result.Data)
}

func searchAll(client *api.Client, query string, formatter *output.Formatter) {
	fmt.Println("=== TRACKS ===")
	searchTracks(client, query, formatter)
	
	fmt.Println("\n=== ALBUMS ===")
	searchAlbums(client, query, formatter)
	
	fmt.Println("\n=== ARTISTS ===")
	searchArtists(client, query, formatter)
	
	fmt.Println("\n=== PLAYLISTS ===")
	searchPlaylists(client, query, formatter)
}

func filterTracksByAlbum(tracks []api.Track, albumName string) []api.Track {
	if albumName == "" {
		return tracks
	}

	var filtered []api.Track
	lowerAlbum := strings.ToLower(albumName)
	
	for _, track := range tracks {
		if exact {
			if strings.ToLower(track.Album.Title) == lowerAlbum {
				filtered = append(filtered, track)
			}
		} else {
			if strings.Contains(strings.ToLower(track.Album.Title), lowerAlbum) {
				filtered = append(filtered, track)
			}
		}
	}
	
	return filtered
}