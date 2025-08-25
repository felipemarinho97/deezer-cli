package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/felipemarinho97/deezer-cli/internal/api"
	"github.com/felipemarinho97/deezer-cli/internal/output"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [type] [id]",
	Short: "Get details for a specific item by ID",
	Long: `Get detailed information for a track, album, artist, playlist, show, or episode by its ID.
	
Examples:
  deezer-cli get track 3135556
  deezer-cli get album 302127 --output json
  deezer-cli get artist 27 --ids-only
  deezer-cli get playlist 908622995
  deezer-cli get show 123456
  deezer-cli get episode 789012`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		itemType := args[0]
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %v\n", err)
			os.Exit(1)
		}

		client := api.NewClient()
		formatter := output.NewFormatter(outputFormat, idsOnly, fields)

		switch itemType {
		case "track":
			getTrack(client, id, formatter)
		case "album":
			getAlbum(client, id, formatter)
		case "artist":
			getArtist(client, id, formatter)
		case "playlist":
			getPlaylist(client, id, formatter)
		case "show", "podcast":
			getShow(client, id, formatter)
		case "episode":
			getEpisode(client, id, formatter)
		default:
			fmt.Fprintf(os.Stderr, "Unknown type: %s. Use track, album, artist, playlist, show, or episode\n", itemType)
			os.Exit(1)
		}
	},
}

var tracksCmd = &cobra.Command{
	Use:   "tracks [type] [id]",
	Short: "Get tracks for an album or top tracks for an artist",
	Long: `Get track listings for albums or top tracks for artists.
	
Examples:
  deezer-cli tracks album 302127
  deezer-cli tracks artist 27 --limit 10
  deezer-cli tracks album 302127 --output json --limit 5`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		itemType := args[0]
		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %v\n", err)
			os.Exit(1)
		}

		client := api.NewClient()
		formatter := output.NewFormatter(outputFormat, idsOnly, fields)

		switch itemType {
		case "album":
			getAlbumTracks(client, id, formatter)
		case "artist":
			getArtistTopTracks(client, id, formatter)
		default:
			fmt.Fprintf(os.Stderr, "Unknown type: %s. Use album or artist\n", itemType)
			os.Exit(1)
		}
	},
}

var albumsCmd = &cobra.Command{
	Use:   "albums artist [id]",
	Short: "Get albums for an artist",
	Long: `Get all albums for a specific artist.
	
Examples:
  deezer-cli albums artist 27
  deezer-cli albums artist 27 --limit 10 --output json`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "artist" {
			fmt.Fprintf(os.Stderr, "Use: deezer-cli albums artist [id]\n")
			os.Exit(1)
		}

		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %v\n", err)
			os.Exit(1)
		}

		client := api.NewClient()
		formatter := output.NewFormatter(outputFormat, idsOnly, fields)
		getArtistAlbums(client, id, formatter)
	},
}

var episodesCmd = &cobra.Command{
	Use:   "episodes show [id]",
	Short: "Get episodes for a podcast show",
	Long: `Get all episodes for a specific podcast show, ordered by most recent.
	
Examples:
  deezer-cli episodes show 406562
  deezer-cli episodes show 406562 --limit 10 --output json`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "show" {
			fmt.Fprintf(os.Stderr, "Use: deezer-cli episodes show [id]\n")
			os.Exit(1)
		}

		id, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %v\n", err)
			os.Exit(1)
		}

		client := api.NewClient()
		formatter := output.NewFormatter(outputFormat, idsOnly, fields)
		getShowEpisodes(client, id, formatter)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(tracksCmd)
	rootCmd.AddCommand(albumsCmd)
	rootCmd.AddCommand(episodesCmd)
}

func getTrack(client *api.Client, id int64, formatter *output.Formatter) {
	track, err := client.GetTrack(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting track: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatTrack(track)
}

func getAlbum(client *api.Client, id int64, formatter *output.Formatter) {
	album, err := client.GetAlbum(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting album: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatAlbum(album)
}

func getArtist(client *api.Client, id int64, formatter *output.Formatter) {
	artist, err := client.GetArtist(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting artist: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatArtist(artist)
}

func getPlaylist(client *api.Client, id int64, formatter *output.Formatter) {
	playlist, err := client.GetPlaylist(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting playlist: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatPlaylist(playlist)
}

func getShow(client *api.Client, id int64, formatter *output.Formatter) {
	show, err := client.GetShow(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting show: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatShow(show)
}

func getEpisode(client *api.Client, id int64, formatter *output.Formatter) {
	episode, err := client.GetEpisode(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting episode: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatEpisode(episode)
}

func getAlbumTracks(client *api.Client, id int64, formatter *output.Formatter) {
	result, err := client.GetAlbumTracks(id, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting album tracks: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatTracks(result.Data)
}

func getArtistTopTracks(client *api.Client, id int64, formatter *output.Formatter) {
	result, err := client.GetArtistTopTracks(id, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting artist top tracks: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatTracks(result.Data)
}

func getArtistAlbums(client *api.Client, id int64, formatter *output.Formatter) {
	result, err := client.GetArtistAlbums(id, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting artist albums: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatAlbums(result.Data)
}

func getShowEpisodes(client *api.Client, id int64, formatter *output.Formatter) {
	result, err := client.GetShowEpisodes(id, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting show episodes: %v\n", err)
		os.Exit(1)
	}

	formatter.FormatEpisodes(result.Data)
}
