package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type StatsMsg struct {
	RamPercent float64
	TotalRam   int
	CpuPercent float64
	TotalCpu   int
}

func getSystemStats() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(500 * time.Millisecond)

		v, _ := mem.VirtualMemory()
		c, _ := cpu.Percent(0, false)
		totalCpu, _ := cpu.Counts(false)

		return StatsMsg{
			RamPercent: v.UsedPercent,
			TotalRam:   int(v.Total),
			CpuPercent: c[0],
			TotalCpu:   totalCpu,
		}
	}
}
