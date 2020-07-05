package usecases

import "gomart-api/master/models"

type ReportUsecase interface {
	GetDailyReport(year, month, day string) ([]*models.Report, error)
	GetMonthlyReport(month string) ([]*models.Report, error)
}
