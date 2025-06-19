package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// shortenCmd represents the shorten command
var shortenCmd = &cobra.Command{
	Use:   "shorten <url>",
	Short: "Shorten a long URL using the BarkURL API",
	Long: `📎 Shorten a URL through your BarkURL server.

Example:
  barkurl shorten https://openai.com
  barkurl --api http://localhost:8080 shorten https://youtube.com

By default, it connects to http://localhost:8080 unless overridden with --api.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		longURL := args[0]

		// Prepare payload
		payload := map[string]string{"url": longURL}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("❌ Failed to create JSON payload:", err)
			os.Exit(1)
		}

		// Send POST request to BarkURL API
		resp, err := http.Post(apiURL+"/shorten", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("❌ Could not reach BarkURL API at", apiURL)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			fmt.Println("❌ API returned error:")
			fmt.Println(string(body))
			os.Exit(1)
		}

		fmt.Println("📎 Shortened URL:")
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
