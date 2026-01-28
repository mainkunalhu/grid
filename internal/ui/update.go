package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.RamBar.Width = msg.Width - 20
		m.CpuBar.Width = msg.Width - 20
		if m.RamBar.Width > MaxWidth {
			m.RamBar.Width = MaxWidth
			m.CpuBar.Width = MaxWidth
		}
		return m, nil

	case StatsMsg:
		m.ramPercent = msg.RamPercent
		m.totalRam = msg.TotalRam
		m.cpuPercent = msg.CpuPercent
		m.totalCpu = msg.TotalCpu

		cmdRam := m.RamBar.SetPercent(m.ramPercent / 100.0)
		cmdCpu := m.CpuBar.SetPercent(m.cpuPercent / 100.0)

		return m, tea.Batch(cmdRam, cmdCpu, getSystemStats())

	case progress.FrameMsg:
		newRam, cmdRam := m.RamBar.Update(msg)
		newCpu, cmdCpu := m.CpuBar.Update(msg)

		m.RamBar = newRam.(progress.Model)
		m.CpuBar = newCpu.(progress.Model)

		return m, tea.Batch(cmdRam, cmdCpu)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}
