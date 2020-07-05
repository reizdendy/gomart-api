package usecases

import (
	"gomart-api/master/models"
	"gomart-api/master/repositories"
)

type reportUsecaseImpl struct {
	reportRepo repositories.ReportRepository
}

func (c reportUsecaseImpl) GetDailyReport(year, month, day string) ([]*models.Report, error) {
	reports, err := c.reportRepo.GetDailyReport(year, month, day)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (c reportUsecaseImpl) GetMonthlyReport(month string) ([]*models.Report, error) {
	reports, err := c.reportRepo.GetMonthlyReport(month)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func InitReportUsecase(reportRepo repositories.ReportRepository) ReportUsecase {
	return &reportUsecaseImpl{reportRepo}
}
