package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// --- Cpu/Ram ---
	ramPercent float64
	totalRam   int
	cpuPercent float64
	totalCpu   int

	// --- Storage Fields ---
	storagePercent float64
	totalStorage   int
	usedStorage    int

	// --- UI Components ---
	RamBar     progress.Model
	CpuBar     progress.Model
	StorageBar progress.Model
}

func toGB(bytes int) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func (m Model) Init() tea.Cmd {
	return getSystemStats()
}
