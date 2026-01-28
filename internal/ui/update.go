package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MaxWidth = 80
	Padding  = 20
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		availableWidth := msg.Width - Padding

		m.RamBar.Width = min(availableWidth, MaxWidth)
		m.CpuBar.Width = min(availableWidth, MaxWidth)
		m.StorageBar.Width = min(availableWidth, MaxWidth)
		return m, nil

	case StatsMsg:
		m.ramPercent = msg.RamPercent
		m.totalRam = msg.TotalRam
		m.cpuPercent = msg.CpuPercent
		m.totalCpu = msg.TotalCpu
		m.storagePercent = msg.StoragePercent
		m.totalStorage = msg.TotalStorage
		m.usedStorage = msg.UsedStorage

		cmdRam := m.RamBar.SetPercent(m.ramPercent / 100.0)
		cmdCpu := m.CpuBar.SetPercent(m.cpuPercent / 100.0)
		cmdStor := m.StorageBar.SetPercent(m.storagePercent / 100.0)

		return m, tea.Batch(cmdRam, cmdCpu, cmdStor, getSystemStats())

	case progress.FrameMsg:
		newRam, cmdRam := m.RamBar.Update(msg)
		newCpu, cmdCpu := m.CpuBar.Update(msg)
		newStor, cmdStor := m.StorageBar.Update(msg)

		m.RamBar = newRam.(progress.Model)
		m.CpuBar = newCpu.(progress.Model)
		m.StorageBar = newStor.(progress.Model)

		return m, tea.Batch(cmdRam, cmdCpu, cmdStor)

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

