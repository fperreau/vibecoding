package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Récupère une tâche par son identifiant",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		resp, err := http.Get(fmt.Sprintf("http://localhost:8080/tasks/%d", id))
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("Tâche non trouvée")
			return
		}

		body, _ := io.ReadAll(resp.Body)
		var task map[string]interface{}
		json.Unmarshal(body, &task)
		fmt.Printf("%d - %s\n", int(task["id"].(float64)), task["description"])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}