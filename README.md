# Deezer CLI

A powerful command-line interface for browsing and searching the Deezer music catalog.

## Features

- **Search**: Find tracks, albums, artists, and playlists
- **Browse**: Get detailed information by ID
- **Multiple Output Formats**: Table (human-readable), JSON, CSV, YAML, IDs-only
- **Advanced Filtering**: Case-insensitive filters for artist and album names
- **Unix-Friendly**: Designed for piping and command chaining
- **Rate Limited**: Respects Deezer API limits
- **Caching**: Optional caching for frequently accessed data

## Installation

```bash
go install github.com/felipemarinho97/deezer-cli@latest
```

Or build from source:

```bash
git clone https://github.com/felipemarinho97/deezer-cli.git
cd deezer-cli
go build -o deezer
```

## Usage

### Search Commands

Search for tracks:
```bash
deezer-cli search "get lucky" --type track
deezer-cli search "daft punk" --type track --limit 10
deezer-cli search "random access" --type track --artist "daft punk"
```

Search for albums:
```bash
deezer-cli search "discovery" --type album
deezer-cli search "homework" --type album --artist "daft punk"
```

Search for artists:
```bash
deezer-cli search "madonna" --type artist
deezer-cli search "queen" --type artist --output json
```

Search for playlists:
```bash
deezer-cli search "workout" --type playlist
deezer-cli search "chill" --type playlist --limit 5
```

### Get Details by ID

Get track details:
```bash
deezer-cli get track 3135556
deezer-cli get track 3135556 --output json
```

Get album details:
```bash
deezer-cli get album 302127
deezer-cli get album 302127 --output yaml
```

Get artist details:
```bash
deezer-cli get artist 27
```

Get playlist details:
```bash
deezer-cli get playlist 908622995
```

### Get Related Content

Get album tracks:
```bash
deezer-cli tracks album 302127
deezer-cli tracks album 302127 --limit 5 --output json
```

Get artist's top tracks:
```bash
deezer-cli tracks artist 27 --limit 10
```

Get artist's albums:
```bash
deezer-cli albums artist 27
deezer-cli albums artist 27 --output csv
```

### Output Formats

Table (default - human readable):
```bash
deezer-cli search "get lucky" --type track
```

JSON (for programmatic use):
```bash
deezer-cli search "get lucky" --type track --output json
```

CSV (for spreadsheets):
```bash
deezer-cli search "get lucky" --type track --output csv
```

YAML:
```bash
deezer-cli search "get lucky" --type track --output yaml
```

IDs only (for piping):
```bash
deezer-cli search "get lucky" --type track --ids-only
```

### Advanced Filtering

Filter by artist (case-insensitive):
```bash
deezer-cli search "lucky" --type track --artist "daft punk"
```

Filter by album:
```bash
deezer-cli search "get" --type track --album "random access"
```

Exact matching:
```bash
deezer-cli search "get" --type track --artist "daft punk" --exact
```

### Piping Examples

Get all track IDs for an artist and fetch details:
```bash
deezer-cli search "daft punk" --type track --ids-only | while read id; do
  deezer-cli get track $id --output json
done
```

Search and filter with jq:
```bash
deezer-cli search "madonna" --type track --output json | jq '.[] | select(.rank > 500000)'
```

Create a playlist of top tracks:
```bash
deezer-cli tracks artist 27 --limit 20 --ids-only > top_tracks.txt
```

## Global Options

- `--output, -o`: Output format (table, json, csv, yaml, ids)
- `--limit, -l`: Limit number of results (default: 25)
- `--ids-only`: Display only IDs
- `--fields, -f`: Select specific fields to display

## Configuration

The CLI can be configured via `~/.config/deezer-cli/config.json`:

```json
{
  "default_format": "table",
  "default_limit": 25,
  "cache_enabled": true,
  "cache_ttl_seconds": 300
}
```

## Examples

### Find all tracks from an album
```bash
album_id=$(deezer-cli search "discovery" --type album --artist "daft punk" --ids-only | head -1)
deezer-cli tracks album $album_id
```

### Export artist discography to CSV
```bash
artist_id=$(deezer-cli search "queen" --type artist --ids-only | head -1)
deezer-cli albums artist $artist_id --output csv > queen_albums.csv
```

### Get random track from search
```bash
deezer-cli search "jazz" --type track --ids-only | shuf -n 1 | xargs -I {} deezer-cli get track {}
```

## API Limits

The CLI respects Deezer API rate limits (20 requests per second) and implements automatic rate limiting.

## License

MIT