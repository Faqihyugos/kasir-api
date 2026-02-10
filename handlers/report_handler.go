package handlers

import (
	"encoding/json"
	"fmt"
	"kasir-api/services"
	"net/http"
	"time"
)

type ReportHandler struct {
	service services.ReportServiceInterface
}

func NewReportHandler(service services.ReportServiceInterface) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleDailySummary godoc
// @Summary Get today's sales summary
// @Description Get sales summary for today
// @Tags reports
// @Accept json
// @Produce json
// @Success 200 {object} models.SalesSummary
// @Failure 500 {string} string "Internal server error"
// @Router /report/hari-ini [get]
func (h *ReportHandler) HandleDailySummary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	summary, err := h.service.GetSalesSummaryToday()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

// HandleDateRangeSummary godoc
// @Summary Get sales summary by date range
// @Description Get sales summary for a specific date range
// @Tags reports
// @Accept json
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} models.SalesSummary
// @Failure 400 {string} string "Invalid date format or missing parameters"
// @Failure 500 {string} string "Internal server error"
// @Router /report [get]
func (h *ReportHandler) HandleDateRangeSummary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	if startDateStr == "" || endDateStr == "" {
		http.Error(w, "start_date and end_date parameters are required", http.StatusBadRequest)
		return
	}

	// Parse dates (format: 2026-01-01)
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid start_date format: %s. Expected format: YYYY-MM-DD", startDateStr), http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid end_date format: %s. Expected format: YYYY-MM-DD", endDateStr), http.StatusBadRequest)
		return
	}

	// Add 1 day to end_date to include the entire end date
	endDate = endDate.Add(24 * time.Hour)

	summary, err := h.service.GetSalesSummaryByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
