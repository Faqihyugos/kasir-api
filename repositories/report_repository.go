package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetSalesSummaryToday mendapatkan summary penjualan hari ini
func (repo *ReportRepository) GetSalesSummaryToday() (*models.SalesSummary, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	return repo.GetSalesSummaryByDateRange(startOfDay, endOfDay)
}

// GetSalesSummaryByDateRange mendapatkan summary penjualan berdasarkan range tanggal
func (repo *ReportRepository) GetSalesSummaryByDateRange(startDate, endDate time.Time) (*models.SalesSummary, error) {
	summary := &models.SalesSummary{}

	// Get total revenue and total transactions
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, startDate, endDate).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)

	if err != nil {
		return nil, err
	}

	// Get top selling product
	var productName sql.NullString
	var qtyTerjual sql.NullInt64

	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, startDate, endDate).Scan(&productName, &qtyTerjual)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Set top product if exists
	if productName.Valid && qtyTerjual.Valid {
		summary.ProdukTerlaris = &models.TopProduct{
			Nama:       productName.String,
			QtyTerjual: int(qtyTerjual.Int64),
		}
	}

	return summary, nil
}
