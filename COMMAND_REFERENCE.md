# Deezer CLI Command Reference

Complete reference for all Deezer CLI commands, flags, and behaviors.

## Global Flags

These flags are available for all commands:

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--output` | `-o` | string | `table` | Output format: table, json, csv, yaml, ids |
| `--limit` | `-l` | int | `25` | Limit number of results |
| `--ids-only` | | boolean | `false` | Display only IDs (overrides --output) |
| `--fields` | `-f` | []string | `[]` | Select specific fields to display |
| `--help` | `-h` | | | Show help for command |

## Commands

### deezer search

Search the Deezer catalog for music content.

**Usage:** `deezer search [query] [flags]`

**Arguments:**
- `query` (required): Search term

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--type` | `-t` | string | `all` | Type to search: track, album, artist, playlist, all |
| `--artist` | | string | `""` | Filter results by artist name (case-insensitive) |
| `--album` | | string | `""` | Filter results by album name (case-insensitive) |
| `--exact` | | boolean | `false` | Use exact matching for filters |

**Behavior:**
- When `--type all`: Shows results in sections (TRACKS, ALBUMS, ARTISTS, PLAYLISTS)
- Artist/album filters work with partial matches unless `--exact` is used
- Results are ranked by relevance/popularity

**Examples:**
```bash
deezer search "daft punk" --type artist
deezer search "get lucky" --artist "daft punk" --exact
deezer search "chill" --type playlist --output json
```

### deezer get

Get detailed information for a specific item by ID.

**Usage:** `deezer get [type] [id] [flags]`

**Arguments:**
- `type` (required): Item type: track, album, artist, playlist
- `id` (required): Numeric ID of the item

**Behavior:**
- Returns detailed information in a formatted view by default
- JSON/YAML outputs include all available fields
- IDs-only mode returns just the ID (useful for validation)

**Examples:**
```bash
deezer get track 3135556
deezer get album 302127 --output json
deezer get artist 27 --ids-only
```

### deezer tracks

Get track listings for albums or top tracks for artists.

**Usage:** `deezer tracks [type] [id] [flags]`

**Arguments:**
- `type` (required): Item type: album, artist
- `id` (required): Numeric ID of the album or artist

**Behavior:**
- For albums: Returns all tracks in the album
- For artists: Returns top/popular tracks
- Results include full track metadata (artist, album, duration, etc.)

**Examples:**
```bash
deezer tracks album 302127
deezer tracks artist 27 --limit 10
deezer tracks album 302127 --output json --limit 5
```

### deezer albums

Get albums for an artist.

**Usage:** `deezer albums artist [id] [flags]`

**Arguments:**
- Must use `artist` as the type (only supported type)
- `id` (required): Numeric ID of the artist

**Behavior:**
- Returns all albums by the specified artist
- Includes album metadata (title, tracks count, release date, etc.)
- Results may include compilations and collaborations

**Examples:**
```bash
deezer albums artist 27
deezer albums artist 27 --limit 10 --output json
```

## Output Formats

### table (default)
- Formatted table with colors and borders
- Optimized for human readability
- Includes truncated text to fit terminal width
- Uses Unicode characters for borders

### json
- Complete JSON data as returned by the API
- Suitable for programmatic processing
- Includes all available fields
- Properly formatted and indented

### csv
- Comma-separated values
- Header row included
- Suitable for spreadsheet applications
- Values are properly escaped

### yaml
- YAML format with proper indentation
- Human-readable structured data
- Good for configuration-like use cases
- Preserves data types

### ids
- One ID per line
- Minimal output for piping to other commands
- Same as using `--ids-only` flag

## Data Fields

### Track Fields
| Field | Description | Type |
|-------|-------------|------|
| ID | Unique track identifier | int64 |
| Title | Song title | string |
| Artist.Name | Primary artist name | string |
| Artist.ID | Artist identifier | int64 |
| Album.Title | Album title | string |
| Album.ID | Album identifier | int64 |
| Duration | Track length in seconds | int |
| Rank | Popularity ranking | int |
| ExplicitLyrics | Contains explicit content | boolean |
| Preview | 30-second preview URL | string |
| Link | Deezer web link | string |

### Album Fields
| Field | Description | Type |
|-------|-------------|------|
| ID | Unique album identifier | int64 |
| Title | Album title | string |
| Artist.Name | Primary artist name | string |
| Artist.ID | Artist identifier | int64 |
| NbTracks | Number of tracks | int |
| ReleaseDate | Release date (YYYY-MM-DD) | string |
| RecordType | Type: album, single, ep | string |
| ExplicitLyrics | Contains explicit content | boolean |
| CoverBig | High-resolution cover URL | string |
| Link | Deezer web link | string |
| Tracklist | API endpoint for tracks | string |

### Artist Fields
| Field | Description | Type |
|-------|-------------|------|
| ID | Unique artist identifier | int64 |
| Name | Artist name | string |
| NbAlbum | Number of albums | int |
| NbFan | Number of fans | int |
| PictureBig | High-resolution photo URL | string |
| Link | Deezer web link | string |
| Tracklist | API endpoint for top tracks | string |

### Playlist Fields
| Field | Description | Type |
|-------|-------------|------|
| ID | Unique playlist identifier | int64 |
| Title | Playlist title | string |
| Description | Playlist description | string |
| Creator.Name | Creator username | string |
| Creator.ID | Creator identifier | int64 |
| NbTracks | Number of tracks | int |
| Duration | Total duration in seconds | int |
| Public | Is publicly visible | boolean |
| Collaborative | Allows collaboration | boolean |
| Fans | Number of followers | int |
| PictureBig | High-resolution cover URL | string |
| Link | Deezer web link | string |

## Error Handling

### Common Errors

| Error | Cause | Solution |
|-------|-------|----------|
| `Invalid ID: strconv.ParseInt: parsing "X": invalid syntax` | Non-numeric ID provided | Use numeric IDs only |
| `Unknown type: X. Use track, album, artist, or playlist` | Invalid type argument | Check command-specific valid types |
| `accepts N arg(s), received M` | Wrong number of arguments | Check command usage with --help |
| `Error getting X: ...` | API error or item not found | Verify ID exists and try again |

### HTTP Errors
- Network timeouts: Check internet connection
- Rate limiting: Wait before making more requests
- API unavailable: Try again later

## Performance Considerations

### Response Times
- Search queries: ~500ms-2s depending on complexity
- Get operations: ~200ms-500ms
- Track/album listings: ~300ms-1s

### Rate Limiting
- No explicit rate limits documented
- Recommended: Max 1 request per second for bulk operations
- Use `--limit` to reduce response size and time

### Memory Usage
- Table format: Low memory usage
- JSON format: Higher memory usage for large result sets
- Streaming not supported - all results loaded into memory

## Integration Examples

### Shell Scripting
```bash
# Get all track IDs from an album
TRACK_IDS=$(deezer tracks album 302127 --ids-only)

# Search and pipe to file
deezer search "jazz" --type playlist --output csv > jazz_playlists.csv

# Check if track exists
if deezer get track 12345 --ids-only >/dev/null 2>&1; then
    echo "Track exists"
fi
```

### Data Processing
```bash
# Convert to different formats
deezer search "rock" --type artist --output json | jq '.[].name'

# Count results
deezer search "pop" --limit 100 --ids-only | wc -l

# Extract specific fields
deezer get artist 27 --output json | jq '.name, .nb_fan'
```

### API Integration
```bash
# Build API request URLs
BASE_URL="https://api.deezer.com"
TRACK_ID=$(deezer search "bohemian rhapsody" --type track --limit 1 --ids-only)
curl "${BASE_URL}/track/${TRACK_ID}"
```

## Troubleshooting

### Common Issues

1. **Command not found**
   - Ensure binary is built: `go build -o deezer .`
   - Check executable permissions: `chmod +x deezer`

2. **No results found**
   - Try broader search terms
   - Check spelling and remove special characters
   - Use different search types

3. **Output formatting issues**
   - Terminal width affects table display
   - Use JSON/CSV for consistent formatting
   - Check terminal color support

4. **Performance problems**
   - Reduce `--limit` value
   - Use more specific search terms
   - Check network connection

### Debug Mode
Currently no debug mode available. For troubleshooting:
- Use `--output json` to see raw API responses
- Check error messages for specific issues
- Verify IDs with smaller test queries

## Version Information

- Go version requirement: 1.21+
- Dependencies: See go.mod for current versions
- API compatibility: Deezer Web API v1.0