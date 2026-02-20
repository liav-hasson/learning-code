package main

import (
	"fmt"
	"os"

	"todo-app/internal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(internal.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
