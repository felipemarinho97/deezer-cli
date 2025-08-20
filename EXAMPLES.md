# Deezer CLI Usage Examples

This document provides comprehensive examples for all Deezer CLI commands.

## Building the Project

```bash
go build -o deezer-cli .
deezer-cli --help
```

## Search Commands

### Basic Search

#### Search All Types
```bash
# Search for "beatles" across all content types
deezer-cli search "beatles" --limit 2

# Output shows tracks, albums, artists, and playlists in separate sections
```

#### Search Specific Types
```bash
# Search for artists
deezer-cli search "daft punk" --type artist --limit 3

# Search for tracks
deezer-cli search "get lucky" --type track --limit 3

# Search for albums
deezer-cli search "random access memories" --type album --limit 5

# Search for playlists
deezer-cli search "chill" --type playlist --limit 3
```

### Advanced Search with Filters

#### Artist Filter
```bash
# Search for tracks by specific artist
deezer-cli search "get lucky" --artist "daft punk" --exact --limit 2

# Case-insensitive artist filtering
deezer-cli search "love" --artist "beatles" --limit 5
```

#### Album Filter
```bash
# Search for tracks from specific album
deezer-cli search "one more time" --album "discovery" --limit 3
```

### Output Formats

#### JSON Output
```bash
deezer-cli search "daft punk" --type artist --limit 2 --output json
```

#### CSV Output
```bash
deezer-cli search "beatles" --type album --limit 2 --output csv
```

#### YAML Output
```bash
deezer-cli search "mozart" --type artist --limit 2 --output yaml
```

#### IDs Only
```bash
deezer-cli search "queen" --type artist --limit 2 --ids-only
```

## Get Commands

### Get Track Details
```bash
# Get specific track by ID
deezer-cli get track 67238735

# Get track in JSON format
deezer-cli get track 67238735 --output json

# Get only track ID
deezer-cli get track 67238735 --ids-only
```

### Get Artist Details
```bash
# Get specific artist by ID
deezer-cli get artist 27

# Get artist in different formats
deezer-cli get artist 27 --output json
deezer-cli get artist 27 --output yaml
```

### Get Album Details
```bash
# Get specific album by ID
deezer-cli get album 302127

# Get album with JSON output
deezer-cli get album 302127 --output json
```

### Get Playlist Details
```bash
# Get specific playlist by ID
deezer-cli get playlist 1311397405

# Get playlist metadata only
deezer-cli get playlist 1311397405 --ids-only
```

## Tracks Commands

### Album Tracks
```bash
# Get all tracks from an album
deezer-cli tracks album 302127 --limit 5

# Get album tracks in JSON format
deezer-cli tracks album 302127 --output json --limit 5

# Get only track IDs from album
deezer-cli tracks album 302127 --ids-only --limit 10
```

### Artist Top Tracks
```bash
# Get top tracks for an artist
deezer-cli tracks artist 27 --limit 5

# Get artist top tracks in CSV format
deezer-cli tracks artist 27 --limit 10 --output csv
```

## Albums Commands

### Artist Albums
```bash
# Get all albums for an artist
deezer-cli albums artist 27 --limit 5

# Get artist albums in JSON format
deezer-cli albums artist 27 --limit 10 --output json

# Get only album IDs
deezer-cli albums artist 27 --ids-only --limit 5
```

## Global Flags

### Limit Results
```bash
# Limit to 3 results
deezer-cli search "rock" --limit 3

# Default limit is 25
deezer-cli search "pop"
```

### Output Formats
All commands support these output formats:
- `table` (default) - Formatted table with colors
- `json` - JSON format for programmatic use
- `csv` - CSV format for spreadsheet import
- `yaml` - YAML format
- `ids` - Only show IDs (same as --ids-only flag)

### Field Selection
```bash
# Select specific fields (behavior varies by output format)
deezer-cli search "queen" --fields title,artist --limit 3
```

## Error Handling Examples

### Invalid Arguments
```bash
# Invalid track ID
deezer-cli get track invalid_id
# Output: Invalid ID: strconv.ParseInt: parsing "invalid_id": invalid syntax

# Invalid type
deezer-cli get invalid_type 123
# Output: Unknown type: invalid_type. Use track, album, artist, or playlist

# Missing search query
deezer-cli search
# Output: Error: accepts 1 arg(s), received 0
```

## Common Use Cases

### Find and Explore an Artist
```bash
# 1. Search for the artist
deezer-cli search "daft punk" --type artist --limit 1

# 2. Get detailed artist information
deezer-cli get artist 27

# 3. Get their top tracks
deezer-cli tracks artist 27 --limit 5

# 4. Get their albums
deezer-cli albums artist 27 --limit 5

# 5. Get tracks from a specific album
deezer-cli tracks album 302127 --limit 5
```

### Export Data for Analysis
```bash
# Export search results to CSV
deezer-cli search "electronic" --type track --limit 100 --output csv > electronic_tracks.csv

# Export artist discography to JSON
deezer-cli albums artist 27 --output json > daft_punk_albums.json

# Get track IDs for further processing
deezer-cli search "house music" --type track --limit 50 --ids-only > house_track_ids.txt
```

### Quick Information Lookup
```bash
# Quick track lookup
deezer-cli get track 67238735 --ids-only

# Quick artist stats
deezer-cli get artist 27 | grep -E "(Name|Albums|Fans)"

# Find specific track by artist and song name
deezer-cli search "get lucky" --artist "daft punk" --exact --limit 1
```

## Performance Notes

- The API has rate limiting, so avoid making too many rapid requests
- Use `--limit` to control the number of results and improve response time
- JSON/CSV outputs are more efficient for large datasets than table format
- Use `--ids-only` when you only need identifiers for further processing

## Tips and Tricks

1. **Exact Matching**: Use `--exact` flag for precise artist/album filtering
2. **Output Piping**: Combine with standard Unix tools for further processing
3. **Batch Processing**: Use `--ids-only` with shell scripts for bulk operations
4. **Data Export**: CSV format works well with spreadsheet applications
5. **API Integration**: JSON output is perfect for integration with other applications