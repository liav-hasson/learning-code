package internal

import "os"

// go's json package can only parse capital letter feilds
type Tasks struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Body     string `json:"body"`
	Priority string `json:"priority"`

	// TODO: implenent task deadline
	// start_date string
	// end_date string
}

func getExistingTaskList() (error, *[]Tasks) {
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		return err, nil
	}

	// TODO: implement file parsing
}

func createNewTaskList() *[]Tasks {
	return &[]Tasks{}
}
