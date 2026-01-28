package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mainkunalhu/grid/internal/ui"
)

func main() {
	ramBar := progress.New(progress.WithDefaultGradient())
	cpuBar := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))

	m := ui.Model{
		RamBar: ramBar,
		CpuBar: cpuBar,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(1)
	}
}

