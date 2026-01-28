package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ram          int
	totalRam     int
	storage      int
	totalStorage int
	cpu          int
	totalCpu     int
}

func InitialModel() Model {
	return Model{
		ram:          0,
		totalRam:     0,
		storage:      0,
		totalStorage: 0,
		cpu:          0,
		totalCpu:     0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
