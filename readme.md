# Inventory Indra — Backend

REST API backend untuk sistem manajemen inventaris, dibangun menggunakan **Go** dengan framework **Gin**. Proyek ini menangani autentikasi pengguna, manajemen produk, stok, transaksi, dan laporan statistik.

---

## Tech Stack

| Layer        | Teknologi                 |
| ------------ | ------------------------- |
| Language     | Go 1.25                   |
| Framework    | Gin                       |
| ORM          | GORM                      |
| Database     | PostgreSQL (via `pgx/v5`) |
| Migration    | Goose                     |
| Auth         | JWT (`golang-jwt/jwt v5`) |
| Excel Export | Excelize v2               |
| Environment  | godotenv                  |
| CORS         | gin-contrib/cors          |

---

## Struktur Direktori

```
inventory-indra-be/
├── cmd/
│   └── migrate/         # Entry point untuk menjalankan database migration
├── db/                  # Koneksi database
├── handler/             # HTTP handler untuk setiap resource
├── helper/              # Fungsi utilitas/helper
├── lib/                 # Library internal
├── middleware/          # Middleware (auth token, handler)
├── model/               # Definisi struct/model data
├── repositories/        # Layer akses data (query ke database)
├── main.go              # Entry point aplikasi
├── go.mod
└── go.sum
```

---

## API Endpoints

Server berjalan di port **`:8080`**.

### Auth & User

| Method  | Endpoint | Middleware      | Deskripsi                             |
| ------- | -------- | --------------- | ------------------------------------- |
| `POST`  | `/user`  | —               | Registrasi pengguna baru              |
| `POST`  | `/login` | —               | Login dan mendapatkan token           |
| `PATCH` | `/user`  | —               | Update data pengguna                  |
| `GET`   | `/user`  | TokenMiddleware | Ambil data pengguna yang sedang login |

### Produk

| Method   | Endpoint   | Middleware        | Deskripsi           |
| -------- | ---------- | ----------------- | ------------------- |
| `POST`   | `/product` | HandlerMiddleware | Tambah produk baru  |
| `GET`    | `/product` | HandlerMiddleware | Ambil daftar produk |
| `PATCH`  | `/product` | HandlerMiddleware | Update data produk  |
| `DELETE` | `/product` | HandlerMiddleware | Hapus produk        |

### Stok

| Method | Endpoint | Middleware        | Deskripsi       |
| ------ | -------- | ----------------- | --------------- |
| `GET`  | `/stock` | HandlerMiddleware | Ambil data stok |

### Transaksi

| Method | Endpoint       | Middleware        | Deskripsi               |
| ------ | -------------- | ----------------- | ----------------------- |
| `POST` | `/transaction` | HandlerMiddleware | Buat transaksi baru     |
| `GET`  | `/transaction` | HandlerMiddleware | Ambil riwayat transaksi |

### Statistik & Laporan

| Method | Endpoint          | Middleware        | Deskripsi                   |
| ------ | ----------------- | ----------------- | --------------------------- |
| `GET`  | `/stats`          | HandlerMiddleware | Ambil data statistik        |
| `GET`  | `/stats/medicine` | HandlerMiddleware | Export data produk ke Excel |
| `GET`  | `/stats/reports`  | HandlerMiddleware | Export laporan ke Excel     |

---

## Konfigurasi CORS

API ini mengizinkan request dari origin berikut:

- `http://localhost:3000` (development)
- `https://inventory-indra.vercel.app` (production)

Allowed methods: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`, `OPTIONS`  
Allowed headers: `Origin`, `Content-Type`, `Authorization`, `Accept`, `x-user-id`

---

## Cara Menjalankan

### Prasyarat

- Go 1.25+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) (untuk migration)

### 1. Clone Repository

```bash
git clone https://github.com/Glenn-Rhee/inventory-indra-be.git
cd inventory-indra-be
```

### 2. Konfigurasi Environment

Buat file `.env` di root direktori:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=inventory_indra
JWT_SECRET=your_jwt_secret
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Jalankan Migration

```bash
go run cmd/migrate/main.go up
```

### 5. Jalankan Server

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`.

---

## Middleware

- **TokenMiddleware** — Memvalidasi JWT token pada header `Authorization`. Digunakan untuk endpoint yang memerlukan autentikasi pengguna biasa.
- **HandlerMiddleware** — Middleware yang lebih umum digunakan untuk sebagian besar endpoint yang memerlukan otorisasi.

---

## Frontend

Proyek ini adalah backend dari aplikasi inventory. Frontend-nya dapat ditemukan di:  
🔗 `https://inventory-indra.vercel.app`

---
