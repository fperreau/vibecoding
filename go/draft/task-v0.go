Gopackage main
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
var tasks = []Task{
	{ID: 1, Description: "Exemple de tâche"},
}
var nextID = 2