package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/spf13/cobra"
)

var (
	apiURL     string
	startRedis bool
)

// rootCmd is the base command
var rootCmd = &cobra.Command{
	Use:   "barkurl",
	Short: "🐾 BarkURL - a friendly URL shortener from the terminal",
	Long: `BarkURL is a CLI and API for shortening URLs using Go, Gin, and Redis.

Examples:
  barkurl shorten https://google.com
  barkurl --api http://localhost:8080 shorten https://openai.com
  barkurl --start-redis shorten https://youtube.com`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if apiURL == "" {
			apiURL = "http://localhost:8080"
		}

		if startRedis {
			fmt.Println("🔁 Attempting to start Redis via Homebrew...")

			// Try starting Redis normally
			startCmd := exec.Command("brew", "services", "start", "redis")
			startCmd.Stdout = os.Stdout
			startCmd.Stderr = os.Stderr

			if err := startCmd.Run(); err != nil {
				panic("⚠️  Failed to start Redis normally.")
			} else {
				fmt.Println("✅ Redis started successfully.")
			}
		}
	},
}

// Execute runs the root CLI command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&apiURL, "api", "a", "http://localhost:8080", "BarkURL API server address")
	rootCmd.PersistentFlags().BoolVar(&startRedis, "start-redis", false, "Start Redis via Homebrew before running")
}
