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

## Kontribusi

Pull request dan issue sangat diterima untuk pengembangan lebih lanjut.

## Lisensi

MIT

## Kontak

- Author: Faqihyugos
  GitHub: [Faqihyugos](https://github.com/Faqihyugos)
