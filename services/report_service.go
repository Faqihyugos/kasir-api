package services

import (
	"kasir-api/models"
	"time"
)

type ReportServiceInterface interface {
	GetSalesSummaryToday() (*models.SalesSummary, error)
	GetSalesSummaryByDateRange(startDate, endDate time.Time) (*models.SalesSummary, error)
}

type ReportService struct {
	repo ReportRepository
}

type ReportRepository interface {
	GetSalesSummaryToday() (*models.SalesSummary, error)
	GetSalesSummaryByDateRange(startDate, endDate time.Time) (*models.SalesSummary, error)
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetSalesSummaryToday() (*models.SalesSummary, error) {
	return s.repo.GetSalesSummaryToday()
}

func (s *ReportService) GetSalesSummaryByDateRange(startDate, endDate time.Time) (*models.SalesSummary, error) {
	return s.repo.GetSalesSummaryByDateRange(startDate, endDate)
}
