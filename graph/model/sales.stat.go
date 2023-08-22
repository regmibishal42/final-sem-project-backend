package model

type MonthlyData struct {
	Month      string  `json:"month"`
	TotalSales float64 `json:"totalSales"`
	TotalUnits int     `json:"totalUnits"`
}

type SalesStatData struct {
	TotalYearlySales      *float64      `json:"totalYearlySales,omitempty"`
	TotalYearlyUnitsSold  *int          `json:"totalYearlyUnitsSold,omitempty"`
	TotalMonthlySales     *float64      `json:"totalMonthlySales,omitempty"`
	TotalMonthlyUnitsSold *int          `json:"totalMonthlyUnitsSold,omitempty"`
	MonthlyData           []MonthlyData `gorm:"-" json:"monthlyData"`
	TotalDailySales       *float64      `json:"totalDailySales,omitempty"`
	TotalDailyUnitsSold   *int          `json:"totalDailyUnitsSold,omitempty"`
}
