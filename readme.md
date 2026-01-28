# Kasir API

RESTful API sederhana untuk aplikasi kasir, dibuat dengan Golang ini menyediakan fitur manajemen produk, category.

## Fitur

- Manajemen produk (CRUD)
- Manajemen category (CRUD)

## Instalasi

1. Pastikan Go sudah terinstall di sistem Anda (versi 1.19 atau lebih baru).
2. Clone repository ini:
   ```bash
   git clone https://github.com/Faqihyugos/kasir-api.git
   cd kasir-api
   ```
3. Install dependencies (jika ada):
   ```bash
   go mod tidy
   ```
4. Jalankan aplikasi:
   ```bash
   go run main.go
   ```

Atau build binary terlebih dahulu:

```bash
go build -o kasir-api main.go
./kasir-api
```

## Penggunaan

API dapat diakses melalui endpoint yang tersedia. Dokumentasi endpoint dapat dilihat pada file dokumentasi atau menggunakan tools seperti Postman.

### API Endpoints

#### Products
- `GET /api/produk` - Get all products (includes category name)
- `GET /api/produk/{id}` - Get product by ID (includes category name)
- `POST /api/produk` - Create new product
- `PUT /api/produk/{id}` - Update product
- `DELETE /api/produk/{id}` - Delete product

#### Categories
- `GET /api/kategori` - Get all categories
- `GET /api/kategori/{id}` - Get category by ID
- `POST /api/kategori` - Create new category
- `PUT /api/kategori/{id}` - Update category
- `DELETE /api/kategori/{id}` - Delete category

### Example Product Response

```json
{
  "id": 1,
  "name": "Laptop Dell XPS",
  "price": 15000000,
  "stock": 10,
  "category_id": 1,
  "category_name": "Electronics"
}
```

## Testing

Jalankan unit tests untuk semua handlers:

```bash
go test ./handlers/... -v
```

Jalankan tests dengan coverage report:

```bash
go test ./handlers/... -cover
```

Lihat detailed coverage:

```bash
go test ./handlers/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Kontribusi

Pull request dan issue sangat diterima untuk pengembangan lebih lanjut.

## Lisensi

MIT

## Kontak

- Author: Faqihyugos
  GitHub: [Faqihyugos](https://github.com/Faqihyugos)
