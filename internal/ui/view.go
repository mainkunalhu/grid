package ui

import (
	"fmt"
)

const (
	Padding  = 2
	MaxWidth = 80
)

func (m Model) View() string {
	usedRamGB := toGB(int(float64(m.totalRam) * (m.ramPercent / 100.0)))
	totalRamGB := toGB(m.totalRam)

	return fmt.Sprintf(
		"\n SYSTEM MONITOR\n\n"+
			" RAM Usage  [%.2f GB / %.2f GB]\n"+
			" %s\n\n"+
			" CPU Usage  [%.2f%% | %d Cores]\n"+
			" %s\n\n"+
			" Press q to quit",
		usedRamGB, totalRamGB,
		m.RamBar.View(),
		m.cpuPercent, m.totalCpu,
		m.CpuBar.View(),
	)
}

