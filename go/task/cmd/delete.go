package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Supprime une tâche par son identifiant",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/tasks/%d", id), nil)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNoContent {
			fmt.Println("Tâche supprimée avec succès")
		} else {
			fmt.Println("Erreur lors de la suppression de la tâche")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}