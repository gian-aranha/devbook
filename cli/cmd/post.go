package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use: "post",
	Short: "Commads to manage posts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand: list, create, etc.")
	},
}

var listPostsCmd = &cobra.Command{
	Use: "list",
	Short: "List all posts",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:5000/posts")
		if err != nil {
			fmt.Println("Error to list posts:", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error to read response body:", err)
			return
		}

		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.AddCommand(listPostsCmd)
}