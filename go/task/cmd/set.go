package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [id] [description]",
	Short: "Met à jour la description d'une tâche",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		description := args[1]
		task := map[string]string{"description": description}
		jsonData, _ := json.Marshal(task)

		client := &http.Client{}
		req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/tasks/%d", id), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Tâche mise à jour avec succès")
		} else {
			fmt.Println("Erreur lors de la mise à jour de la tâche")
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
