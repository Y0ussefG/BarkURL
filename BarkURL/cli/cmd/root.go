package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiURL string // 🐾 shared across all subcommands

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "barkurl",
	Short: "🐾 BarkURL - a friendly URL shortener from the terminal",
	Long:  `BarkURL is a CLI and API for shortening URLs using Go, Gin, and Redis.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// This will run before any subcommand
		if apiURL == "" {
			apiURL = "http://localhost:8080"
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// ✅ Add persistent --api / -a flag to all subcommands
	rootCmd.PersistentFlags().StringVarP(&apiURL, "api", "a", "http://localhost:8080", "BarkURL API server address")
}
