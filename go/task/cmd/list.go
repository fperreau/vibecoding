package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste toutes les tâches",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/tasks")
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var tasks []map[string]interface{}
		json.Unmarshal(body, &tasks)

		for _, task := range tasks {
			fmt.Printf("%d - %s\n", int(task["id"].(float64)), task["description"])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}