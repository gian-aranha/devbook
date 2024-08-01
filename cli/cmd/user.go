package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use: "user",
	Short: "Commads to manage users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand: list, create, etc.")
	},
}

var listUsersCmd = &cobra.Command{
	Use: "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:5000/users")
		if err != nil {
			fmt.Println("Error to list users:", err)
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
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(listUsersCmd)
}