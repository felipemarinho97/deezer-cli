package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/felipemarinho97/deezer-cli/internal/api"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

type Formatter struct {
	format   string
	idsOnly  bool
	fields   []string
}

func NewFormatter(format string, idsOnly bool, fields []string) *Formatter {
	if idsOnly {
		format = "ids"
	}
	return &Formatter{
		format:  format,
		idsOnly: idsOnly,
		fields:  fields,
	}
}

func (f *Formatter) FormatTracks(tracks []api.Track) {
	if len(tracks) == 0 {
		fmt.Println("No tracks found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(tracks)
	case "csv":
		f.outputTracksCSV(tracks)
	case "yaml":
		f.outputYAML(tracks)
	case "ids":
		f.outputTrackIDs(tracks)
	default:
		f.outputTracksTable(tracks)
	}
}

func (f *Formatter) FormatAlbums(albums []api.Album) {
	if len(albums) == 0 {
		fmt.Println("No albums found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(albums)
	case "csv":
		f.outputAlbumsCSV(albums)
	case "yaml":
		f.outputYAML(albums)
	case "ids":
		f.outputAlbumIDs(albums)
	default:
		f.outputAlbumsTable(albums)
	}
}

func (f *Formatter) FormatArtists(artists []api.Artist) {
	if len(artists) == 0 {
		fmt.Println("No artists found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(artists)
	case "csv":
		f.outputArtistsCSV(artists)
	case "yaml":
		f.outputYAML(artists)
	case "ids":
		f.outputArtistIDs(artists)
	default:
		f.outputArtistsTable(artists)
	}
}

func (f *Formatter) FormatPlaylists(playlists []api.Playlist) {
	if len(playlists) == 0 {
		fmt.Println("No playlists found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(playlists)
	case "csv":
		f.outputPlaylistsCSV(playlists)
	case "yaml":
		f.outputYAML(playlists)
	case "ids":
		f.outputPlaylistIDs(playlists)
	default:
		f.outputPlaylistsTable(playlists)
	}
}

func (f *Formatter) FormatShows(shows []api.Show) {
	if len(shows) == 0 {
		fmt.Println("No shows found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(shows)
	case "csv":
		f.outputShowsCSV(shows)
	case "yaml":
		f.outputYAML(shows)
	case "ids":
		f.outputShowIDs(shows)
	default:
		f.outputShowsTable(shows)
	}
}

func (f *Formatter) FormatEpisodes(episodes []api.Episode) {
	if len(episodes) == 0 {
		fmt.Println("No episodes found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(episodes)
	case "csv":
		f.outputEpisodesCSV(episodes)
	case "yaml":
		f.outputYAML(episodes)
	case "ids":
		f.outputEpisodeIDs(episodes)
	default:
		f.outputEpisodesTable(episodes)
	}
}

func (f *Formatter) outputTracksTable(tracks []api.Track) {
	table := tablewriter.NewWriter(os.Stdout)
	
	headers := []string{"ID", "Title", "Artist", "Album", "Duration", "Link"}
	if f.shouldIncludeField("rank", "all") {
		headers = append(headers, "Rank")
	}
	
	table.SetHeader(headers)
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	headerColors := make([]tablewriter.Colors, len(headers))
	for i := range headerColors {
		headerColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor}
	}
	table.SetHeaderColor(headerColors...)

	for _, track := range tracks {
		row := []string{
			strconv.FormatInt(track.ID, 10),
			truncate(track.Title, 30),
			truncate(track.Artist.Name, 20),
			truncate(track.Album.Title, 25),
			track.GetDurationFormatted(),
			track.Link,
		}
		
		if f.shouldIncludeField("rank", "all") {
			row = append(row, strconv.Itoa(track.Rank))
		}
		
		table.Append(row)
	}
	
	table.Render()
}

func (f *Formatter) outputAlbumsTable(albums []api.Album) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Artist", "Tracks", "Release", "Link"})
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)

	for _, album := range albums {
		table.Append([]string{
			strconv.FormatInt(album.ID, 10),
			truncate(album.Title, 30),
			truncate(album.Artist.Name, 20),
			strconv.Itoa(album.NbTracks),
			album.ReleaseDate,
			album.Link,
		})
	}
	
	table.Render()
}

func (f *Formatter) outputArtistsTable(artists []api.Artist) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Albums", "Fans", "Link"})
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
	)

	for _, artist := range artists {
		table.Append([]string{
			strconv.FormatInt(artist.ID, 10),
			truncate(artist.Name, 30),
			strconv.Itoa(artist.NbAlbum),
			formatNumber(artist.NbFan),
			artist.Link,
		})
	}
	
	table.Render()
}

func (f *Formatter) outputPlaylistsTable(playlists []api.Playlist) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Creator", "Tracks", "Public", "Link"})
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
	)

	for _, playlist := range playlists {
		public := "No"
		if playlist.Public {
			public = "Yes"
		}
		
		table.Append([]string{
			strconv.FormatInt(playlist.ID, 10),
			truncate(playlist.Title, 30),
			truncate(playlist.GetCreatorName(), 20),
			strconv.Itoa(playlist.NbTracks),
			public,
			playlist.Link,
		})
	}
	
	table.Render()
}

func (f *Formatter) outputShowsTable(shows []api.Show) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Description", "Available", "Fans", "Link"})
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
	)

	for _, show := range shows {
		available := "No"
		if show.Available {
			available = "Yes"
		}
		
		table.Append([]string{
			strconv.FormatInt(show.ID, 10),
			truncate(show.Title, 30),
			truncate(show.Description, 40),
			available,
			formatNumber(show.Fans),
			show.Link,
		})
	}
	
	table.Render()
}

func (f *Formatter) outputEpisodesTable(episodes []api.Episode) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Show", "Duration", "Release Date", "Available", "Link"})
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetCenterSeparator("│")
	table.SetColumnSeparator("│")
	table.SetRowSeparator("─")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
	)

	for _, episode := range episodes {
		available := "No"
		if episode.Available {
			available = "Yes"
		}
		
		table.Append([]string{
			strconv.FormatInt(episode.ID, 10),
			truncate(episode.Title, 30),
			truncate(episode.Show.Title, 20),
			episode.GetDurationFormatted(),
			episode.ReleaseDate,
			available,
			episode.Link,
		})
	}
	
	table.Render()
}

func (f *Formatter) outputJSON(data interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}

func (f *Formatter) outputYAML(data interface{}) {
	encoder := yaml.NewEncoder(os.Stdout)
	encoder.SetIndent(2)
	encoder.Encode(data)
}

func (f *Formatter) outputTracksCSV(tracks []api.Track) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Artist", "Album", "Duration", "Link", "Rank"})
	
	for _, track := range tracks {
		writer.Write([]string{
			strconv.FormatInt(track.ID, 10),
			track.Title,
			track.Artist.Name,
			track.Album.Title,
			strconv.Itoa(track.Duration),
			track.Link,
			strconv.Itoa(track.Rank),
		})
	}
}

func (f *Formatter) outputAlbumsCSV(albums []api.Album) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Artist", "Tracks", "ReleaseDate", "Link"})
	
	for _, album := range albums {
		writer.Write([]string{
			strconv.FormatInt(album.ID, 10),
			album.Title,
			album.Artist.Name,
			strconv.Itoa(album.NbTracks),
			album.ReleaseDate,
			album.Link,
		})
	}
}

func (f *Formatter) outputArtistsCSV(artists []api.Artist) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Albums", "Fans", "Link"})
	
	for _, artist := range artists {
		writer.Write([]string{
			strconv.FormatInt(artist.ID, 10),
			artist.Name,
			strconv.Itoa(artist.NbAlbum),
			strconv.Itoa(artist.NbFan),
			artist.Link,
		})
	}
}

func (f *Formatter) outputPlaylistsCSV(playlists []api.Playlist) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Creator", "Tracks", "Public", "Link"})
	
	for _, playlist := range playlists {
		writer.Write([]string{
			strconv.FormatInt(playlist.ID, 10),
			playlist.Title,
			playlist.GetCreatorName(),
			strconv.Itoa(playlist.NbTracks),
			strconv.FormatBool(playlist.Public),
			playlist.Link,
		})
	}
}

func (f *Formatter) outputShowsCSV(shows []api.Show) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Description", "Available", "Fans", "Link"})
	
	for _, show := range shows {
		writer.Write([]string{
			strconv.FormatInt(show.ID, 10),
			show.Title,
			show.Description,
			strconv.FormatBool(show.Available),
			strconv.Itoa(show.Fans),
			show.Link,
		})
	}
}

func (f *Formatter) outputEpisodesCSV(episodes []api.Episode) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "Show", "Duration", "ReleaseDate", "Available", "Link"})
	
	for _, episode := range episodes {
		writer.Write([]string{
			strconv.FormatInt(episode.ID, 10),
			episode.Title,
			episode.Show.Title,
			episode.GetDurationFormatted(),
			episode.ReleaseDate,
			strconv.FormatBool(episode.Available),
			episode.Link,
		})
	}
}

func (f *Formatter) outputTrackIDs(tracks []api.Track) {
	for _, track := range tracks {
		fmt.Println(track.ID)
	}
}

func (f *Formatter) outputAlbumIDs(albums []api.Album) {
	for _, album := range albums {
		fmt.Println(album.ID)
	}
}

func (f *Formatter) outputArtistIDs(artists []api.Artist) {
	for _, artist := range artists {
		fmt.Println(artist.ID)
	}
}

func (f *Formatter) outputPlaylistIDs(playlists []api.Playlist) {
	for _, playlist := range playlists {
		fmt.Println(playlist.ID)
	}
}

func (f *Formatter) outputShowIDs(shows []api.Show) {
	for _, show := range shows {
		fmt.Println(show.ID)
	}
}

func (f *Formatter) outputEpisodeIDs(episodes []api.Episode) {
	for _, episode := range episodes {
		fmt.Println(episode.ID)
	}
}

func (f *Formatter) shouldIncludeField(field string, defaultFields ...string) bool {
	if len(f.fields) == 0 {
		for _, df := range defaultFields {
			if df == field || df == "all" {
				return true
			}
		}
		return false
	}
	
	for _, f := range f.fields {
		if strings.ToLower(f) == strings.ToLower(field) {
			return true
		}
	}
	return false
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func formatNumber(n int) string {
	if n < 1000 {
		return strconv.Itoa(n)
	} else if n < 1000000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	} else {
		return fmt.Sprintf("%.1fM", float64(n)/1000000)
	}
}

func (f *Formatter) FormatTrack(track *api.Track) {
	if track == nil {
		fmt.Println("Track not found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(track)
	case "yaml":
		f.outputYAML(track)
	case "ids":
		fmt.Println(track.ID)
	default:
		f.outputTrackDetail(track)
	}
}

func (f *Formatter) FormatAlbum(album *api.Album) {
	if album == nil {
		fmt.Println("Album not found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(album)
	case "yaml":
		f.outputYAML(album)
	case "ids":
		fmt.Println(album.ID)
	default:
		f.outputAlbumDetail(album)
	}
}

func (f *Formatter) FormatArtist(artist *api.Artist) {
	if artist == nil {
		fmt.Println("Artist not found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(artist)
	case "yaml":
		f.outputYAML(artist)
	case "ids":
		fmt.Println(artist.ID)
	default:
		f.outputArtistDetail(artist)
	}
}

func (f *Formatter) FormatPlaylist(playlist *api.Playlist) {
	if playlist == nil {
		fmt.Println("Playlist not found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(playlist)
	case "yaml":
		f.outputYAML(playlist)
	case "ids":
		fmt.Println(playlist.ID)
	default:
		f.outputPlaylistDetail(playlist)
	}
}

func (f *Formatter) FormatShow(show *api.Show) {
	if show == nil {
		fmt.Println("Show not found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(show)
	case "yaml":
		f.outputYAML(show)
	case "ids":
		fmt.Println(show.ID)
	default:
		f.outputShowDetail(show)
	}
}

func (f *Formatter) FormatEpisode(episode *api.Episode) {
	if episode == nil {
		fmt.Println("Episode not found")
		return
	}

	switch f.format {
	case "json":
		f.outputJSON(episode)
	case "yaml":
		f.outputYAML(episode)
	case "ids":
		fmt.Println(episode.ID)
	default:
		f.outputEpisodeDetail(episode)
	}
}

func (f *Formatter) outputTrackDetail(track *api.Track) {
	bold := color.New(color.Bold)
	cyan := color.New(color.FgCyan)
	
	bold.Println("Track Details")
	fmt.Println(strings.Repeat("─", 50))
	
	cyan.Print("ID: ")
	fmt.Println(track.ID)
	
	cyan.Print("Title: ")
	fmt.Println(track.Title)
	
	cyan.Print("Artist: ")
	fmt.Printf("%s (ID: %d)\n", track.Artist.Name, track.Artist.ID)
	
	cyan.Print("Album: ")
	fmt.Printf("%s (ID: %d)\n", track.Album.Title, track.Album.ID)
	
	cyan.Print("Duration: ")
	fmt.Println(track.GetDurationFormatted())
	
	cyan.Print("Rank: ")
	fmt.Println(track.Rank)
	
	cyan.Print("Explicit: ")
	fmt.Println(track.ExplicitLyrics)
	
	cyan.Print("Preview: ")
	fmt.Println(track.Preview)
	
	cyan.Print("Link: ")
	fmt.Println(track.Link)
}

func (f *Formatter) outputAlbumDetail(album *api.Album) {
	bold := color.New(color.Bold)
	green := color.New(color.FgGreen)
	
	bold.Println("Album Details")
	fmt.Println(strings.Repeat("─", 50))
	
	green.Print("ID: ")
	fmt.Println(album.ID)
	
	green.Print("Title: ")
	fmt.Println(album.Title)
	
	green.Print("Artist: ")
	fmt.Printf("%s (ID: %d)\n", album.Artist.Name, album.Artist.ID)
	
	green.Print("Tracks: ")
	fmt.Println(album.NbTracks)
	
	green.Print("Release Date: ")
	fmt.Println(album.ReleaseDate)
	
	green.Print("Record Type: ")
	fmt.Println(album.RecordType)
	
	green.Print("Explicit: ")
	fmt.Println(album.ExplicitLyrics)
	
	green.Print("Cover: ")
	fmt.Println(album.CoverBig)
	
	green.Print("Link: ")
	fmt.Println(album.Link)
	
	green.Print("Tracklist: ")
	fmt.Println(album.Tracklist)
}

func (f *Formatter) outputArtistDetail(artist *api.Artist) {
	bold := color.New(color.Bold)
	yellow := color.New(color.FgYellow)
	
	bold.Println("Artist Details")
	fmt.Println(strings.Repeat("─", 50))
	
	yellow.Print("ID: ")
	fmt.Println(artist.ID)
	
	yellow.Print("Name: ")
	fmt.Println(artist.Name)
	
	yellow.Print("Albums: ")
	fmt.Println(artist.NbAlbum)
	
	yellow.Print("Fans: ")
	fmt.Println(formatNumber(artist.NbFan))
	
	yellow.Print("Picture: ")
	fmt.Println(artist.PictureBig)
	
	yellow.Print("Link: ")
	fmt.Println(artist.Link)
	
	yellow.Print("Tracklist: ")
	fmt.Println(artist.Tracklist)
}

func (f *Formatter) outputPlaylistDetail(playlist *api.Playlist) {
	bold := color.New(color.Bold)
	magenta := color.New(color.FgMagenta)
	
	bold.Println("Playlist Details")
	fmt.Println(strings.Repeat("─", 50))
	
	magenta.Print("ID: ")
	fmt.Println(playlist.ID)
	
	magenta.Print("Title: ")
	fmt.Println(playlist.Title)
	
	magenta.Print("Description: ")
	fmt.Println(playlist.Description)
	
	magenta.Print("Creator: ")
	if playlist.Creator != nil {
		fmt.Printf("%s (ID: %d)\n", playlist.Creator.Name, playlist.Creator.ID)
	} else {
		fmt.Println("Unknown")
	}
	
	magenta.Print("Tracks: ")
	fmt.Println(playlist.NbTracks)
	
	magenta.Print("Duration: ")
	duration := time.Duration(playlist.Duration) * time.Second
	fmt.Printf("%d:%02d:%02d\n", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
	
	magenta.Print("Public: ")
	fmt.Println(playlist.Public)
	
	magenta.Print("Collaborative: ")
	fmt.Println(playlist.Collaborative)
	
	magenta.Print("Fans: ")
	fmt.Println(formatNumber(playlist.Fans))
	
	magenta.Print("Picture: ")
	fmt.Println(playlist.PictureBig)
	
	magenta.Print("Link: ")
	fmt.Println(playlist.Link)
}

func (f *Formatter) outputShowDetail(show *api.Show) {
	bold := color.New(color.Bold)
	magenta := color.New(color.FgMagenta)
	
	bold.Println("Show Details")
	fmt.Println(strings.Repeat("─", 50))
	
	magenta.Print("ID: ")
	fmt.Println(show.ID)
	
	magenta.Print("Title: ")
	fmt.Println(show.Title)
	
	magenta.Print("Description: ")
	fmt.Println(show.Description)
	
	magenta.Print("Available: ")
	fmt.Println(show.Available)
	
	magenta.Print("Fans: ")
	fmt.Println(formatNumber(show.Fans))
	
	magenta.Print("Picture: ")
	fmt.Println(show.PictureBig)
	
	magenta.Print("Link: ")
	fmt.Println(show.Link)
}

func (f *Formatter) outputEpisodeDetail(episode *api.Episode) {
	bold := color.New(color.Bold)
	magenta := color.New(color.FgMagenta)
	
	bold.Println("Episode Details")
	fmt.Println(strings.Repeat("─", 50))
	
	magenta.Print("ID: ")
	fmt.Println(episode.ID)
	
	magenta.Print("Title: ")
	fmt.Println(episode.Title)
	
	magenta.Print("Show: ")
	fmt.Printf("%s (ID: %d)\n", episode.Show.Title, episode.Show.ID)
	
	magenta.Print("Description: ")
	fmt.Println(episode.Description)
	
	magenta.Print("Duration: ")
	fmt.Println(episode.GetDurationFormatted())
	
	magenta.Print("Release Date: ")
	fmt.Println(episode.ReleaseDate)
	
	magenta.Print("Available: ")
	fmt.Println(episode.Available)
	
	magenta.Print("Picture: ")
	fmt.Println(episode.PictureBig)
	
	magenta.Print("Link: ")
	fmt.Println(episode.Link)
}