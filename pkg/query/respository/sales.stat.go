package query_repository

import (
	"backend/graph/model"
	"context"
	"fmt"
	"time"
)

func (r QueryRepository) GetSalesStat(ctx context.Context, input *model.SalesStatInput, organizationID *string) (*model.SalesStatData, error) {
	var salesStat model.SalesStatData

	db := r.db.Model(&model.Sales{}).Where("organization_id = ?", organizationID)
	// Calculate total yearly sales and total yearly units sold
	currentYear := time.Now().Year()
	if err := db.Select("SUM(sold_at) as total_yearly_sales, SUM(units_sold) as total_yearly_units_sold").
		Where("EXTRACT(year FROM created_at) = ?", currentYear).
		Scan(&salesStat).Error; err != nil {
		return nil, err
	}

	// Calculate total monthly sales and total monthly units sold
	currentMonth := time.Now().Month()
	if err := db.Select("SUM(sold_at) as total_monthly_sales, SUM(units_sold) as total_monthly_units_sold").
		Where("EXTRACT(year FROM created_at) = ? AND EXTRACT(month FROM created_at) = ?", currentYear, currentMonth).
		Scan(&salesStat).Error; err != nil {
		return nil, err
	}

	// Calculate total daily sales and total daily units sold
	currentDate := time.Now().Format("2006-01-02")
	if err := db.Select("SUM(sold_at) as total_daily_sales, SUM(units_sold) as total_daily_units_sold").Where("DATE(created_at) = ?", currentDate).
		Scan(&salesStat).Error; err != nil {
		return nil, err
	}
	// Calculate total weekly sales and total weekly units sold
	// _, currentWeek := time.Now().ISOWeek()
	// if err := db.Table("sales").
	// 	Select("SUM(sold_at) as total_weekly_sales, SUM(units_sold) as total_weekly_units_sold").
	// 	Where("EXTRACT(year FROM created_at) = ? AND EXTRACT(week FROM created_at) = ?", currentYear, currentWeek).
	// 	Scan(&salesStat).Error; err != nil {
	// 	log.Printf("Error calculating weekly sales: %v", err)
	// }
	var monthlyData []model.MonthlyData

	query := fmt.Sprintf(`
    SELECT TO_CHAR(DATE_TRUNC('month', created_at), 'Month') as month,
           SUM(sold_at) as total_sales,
           SUM(units_sold) as total_units
    FROM sales
    WHERE EXTRACT(year FROM created_at) = %d
    GROUP BY DATE_TRUNC('month', created_at)
    ORDER BY DATE_TRUNC('month', created_at)
`, currentYear)

	if err := db.Raw(query).Scan(&monthlyData).Error; err != nil {
		return nil, err
	}
	salesStat.MonthlyData = monthlyData
	return &salesStat, nil

}

func (r QueryRepository) GetDailySalesStat(ctx context.Context, organizationID *string) ([]*model.DailySalesData, error) {
	dailyData := []*model.DailySalesData{}
	currentYear := time.Now().Year()
	db := r.db.Model(&model.Sales{}).Where("deleted_at IS NULL AND organization_id = ?", organizationID)
	query := fmt.Sprintf(`
    SELECT DATE(created_at) as date, SUM(sold_at) as total_sales, SUM(units_sold) as total_units
    FROM sales
    WHERE EXTRACT(year FROM created_at) = %d
    GROUP BY DATE(created_at)
    ORDER BY DATE(created_at)
`, currentYear)

	if err := db.Raw(query).Scan(&dailyData).Error; err != nil {
		return nil, err
	}
	return dailyData, nil
}

func (r QueryRepository) GetSalesStatBreakdownByCategory(ctx context.Context, organizationID string, input *model.SalesBreakDownInput) ([]*model.SalesBreakdownData, error) {
	data := []*model.SalesBreakdownData{}

	db := r.db.Model(&model.Sales{}).Where("sales.deleted_at IS NULL AND sales.organization_id = ?", organizationID)
	db = db.Select("SUM(sales.sold_at) as total_sales,categories.name as category_name").Joins("left join products on products.id = sales.product_id").
		Joins("left join categories on products.category_id = categories.id").Group("categories.name")
	err := db.Scan(&data).Error
	if err != nil {
		fmt.Println("Query Error", err)
		return nil, err
	}

	return data, nil
}
