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
	storBar := progress.New(progress.WithScaledGradient("#448AFF", "#00E5FF"))

	m := ui.Model{
		RamBar:     ramBar,
		CpuBar:     cpuBar,
		StorageBar: storBar,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(1)
	}
}

