package repositories

import "gomart-api/master/models"

type ReportRepository interface {
	GetDailyReport(year, month, day string) ([]*models.Report, error)
	GetMonthlyReport(month string) ([]*models.Report, error)
}
