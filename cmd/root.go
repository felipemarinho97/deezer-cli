package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	limit        int
	idsOnly      bool
	fields       []string
)

var rootCmd = &cobra.Command{
	Use:   "deezer",
	Short: "A powerful CLI for browsing the Deezer music catalog",
	Long: `Deezer CLI - Browse and search the Deezer music catalog with ease.
	
This tool allows you to search for tracks, albums, artists, and playlists,
get detailed information by ID, and format output for both human reading
and piping to other commands.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format: table, json, csv, yaml, ids")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 25, "Limit number of results")
	rootCmd.PersistentFlags().BoolVar(&idsOnly, "ids-only", false, "Display only IDs")
	rootCmd.PersistentFlags().StringSliceVarP(&fields, "fields", "f", []string{}, "Select specific fields to display")
}