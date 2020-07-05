package repositories

import (
	"database/sql"
	"fmt"
	"gomart-api/master/models"
	"gomart-api/utils"
)

type reportRepoImpl struct {
	db *sql.DB
}

func (c *reportRepoImpl) GetDailyReport(year, month, day string) ([]*models.Report, error) {
	rows, err := c.db.Query(utils.GET_DAILY_REPORT, year, month, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.Report
	for rows.Next() {
		report := models.Report{}
		err := rows.Scan(&report.TransactionDate, &report.ProductName, &report.ProductPrice, &report.Quantity, &report.Total)

		if err != nil {
			return nil, err
		}

		reports = append(reports, &report)
	}
	fmt.Println(reports)
	return reports, nil
}

func (c *reportRepoImpl) GetMonthlyReport(month string) ([]*models.Report, error) {
	rows, err := c.db.Query(utils.GET_MONTHLY_REPORT, month)
	fmt.Println(month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.Report
	for rows.Next() {
		report := models.Report{}
		err := rows.Scan(&report.TransactionDate, &report.ProductName, &report.ProductPrice, &report.Quantity, &report.Total)

		if err != nil {
			return nil, err
		}

		reports = append(reports, &report)
	}
	return reports, nil
}

func InitReportRepositoryImpl(db *sql.DB) ReportRepository {
	return &reportRepoImpl{db}
}
