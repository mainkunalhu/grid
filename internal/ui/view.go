package ui

import (
	"fmt"
)

func (m Model) View() string {

	ramUsed := toGB(int(float64(m.totalRam) * (m.ramPercent / 100.0)))
	ramTotal := toGB(int(m.totalRam))

	storUsed := toGB(m.usedStorage)
	storTotal := toGB(m.totalStorage)

	return fmt.Sprintf(
		"\n GRID SYSTEM MONITOR\n\n"+
			// RAM
			" RAM Usage     [%.2f GB / %.2f GB]\n"+
			" %s\n\n"+
			// CPU
			" CPU Usage     [%.2f%% | %d Cores]\n"+
			" %s\n\n"+
			// STORAGE
			" Disk Usage    [%.2f GB / %.2f GB]\n"+
			" %s\n\n"+
			" Press q to quit",

		ramUsed, ramTotal, m.RamBar.View(),
		m.cpuPercent, m.totalCpu, m.CpuBar.View(),
		storUsed, storTotal, m.StorageBar.View(),
	)
}
