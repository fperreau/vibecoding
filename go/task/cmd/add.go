package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Ajoute une nouvelle tâche",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		task := map[string]string{"description": description}
		jsonData, _ := json.Marshal(task)

		resp, err := http.Post("http://localhost:8080/tasks", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			fmt.Println("Tâche ajoutée avec succès")
		} else {
			fmt.Println("Erreur lors de l'ajout de la tâche")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}