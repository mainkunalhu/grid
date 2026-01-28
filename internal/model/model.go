package model

type StatsModel struct {
	ram          int
	totalRam     int
	storage      int
	totalStorage int
	cpu          int
	totalCpu     int
}

func InitialModel() StatsModel {
	return StatsModel{
		ram:          0,
		totalRam:     0,
		storage:      0,
		totalStorage: 0,
		cpu:          0,
		totalCpu:     0,
	}
}
