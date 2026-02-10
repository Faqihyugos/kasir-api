package models

type SalesSummary struct {
	TotalRevenue   int         `json:"total_revenue"`
	TotalTransaksi int         `json:"total_transaksi"`
	ProdukTerlaris *TopProduct `json:"produk_terlaris"`
}

type TopProduct struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
