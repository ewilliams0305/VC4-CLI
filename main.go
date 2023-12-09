package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ewilliams0305/VC4-CLI/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("VC4 CLI failed to start, there's been an error: %v", err)
		os.Exit(1)
	}
}
