

# 🚀 GoLang Take Home Test - Aplikasi Data Mahasiswa (REST API & gRPC)

Aplikasi ini mengimplementasikan REST API untuk pengelolaan data Mahasiswa, Jurusan, dan Hobi menggunakan GoLang dan database PostgreSQL. [cite_start]Aplikasi dirancang untuk menjalankan server REST dan gRPC secara bersamaan (concurrent) dalam satu file `main.go`, sesuai dengan aturan tugas[cite: 12].

## 🎯 Persyaratan Tugas yang Dipenuhi

* [cite_start]**Bahasa:** GoLang [cite: 11]
* [cite_start]**Database:** PostgreSQL [cite: 2]
* [cite_start]**Endpoint REST:** Insert data mahasiswa, Get semua data mahasiswa, dan Get detail data mahasiswa by ID[cite: 3, 4, 5].
* [cite_start]**Arsitektur:** Menggunakan REST dan gRPC dalam satu file `main.go`[cite: 12].
* [cite_start]**Unit Test:** Wajib disediakan.
* [cite_start]**Dokumentasi:** File `README.md` ini dan dokumentasi REST API (Postman Collection).

## ⚙️ Setup Environment

### 1. Kebutuhan Sistem

* Go Language (Versi 1.18+)
* PostgreSQL (Versi 10+)
* Klien Git dan Postman/cURL.

### 2. Dependencies Go

Pastikan semua *library* terinstal, termasuk `godotenv` untuk membaca ENV, `gin` untuk REST, dan *package* `grpc`:

go mod tidy
go get [github.com/joho/godotenv](https://github.com/joho/godotenv)

### 3. Dependencies Go

Kredensial database dan port diatur melalui Environment Variables, dibaca dari file .env.

Buat file .env di root direktori proyek dan isi dengan konfigurasi koneksi database Anda:

### 4. SKEMA DATABASE
-- Tabel Jurusan
-- Kolom: ID (Primary Key), Nama (Varchar)
CREATE TABLE jurusan (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL
    );

-- Tabel Hobi
-- Kolom: ID (Primary Key), Nama (Varchar)
CREATE TABLE hobi (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL
    );

-- Tabel Mahasiswa
-- Kolom: ID (Primary Key), Nama (Varchar), Tanggal Lahir (Date), Gender (int)
CREATE TABLE mahasiswa (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    gender INT NOT NULL -- 0: Laki-laki, 1: Perempuan 
    );

-- Tabel Penghubung Mahasiswa dan Jurusan (Mahasiswa Jurusan)
CREATE TABLE mahasiswa_jurusan (
    id SERIAL PRIMARY KEY,
    id_mahasiswa INT NOT NULL REFERENCES mahasiswa(id) ON DELETE CASCADE, -- ID Mahasiswa
    id_jurusan INT NOT NULL REFERENCES jurusan(id) ON DELETE CASCADE,     -- ID Jurusan
    UNIQUE (id_mahasiswa, id_jurusan)
    );

-- Tabel Penghubung Mahasiswa dan Hobi (Mahasiswa Hobi)
CREATE TABLE mahasiswa_hobi (
    id SERIAL PRIMARY KEY,
    id_mahasiswa INT NOT NULL REFERENCES mahasiswa(id) ON DELETE CASCADE, -- ID Mahasiswa
    id_hobi INT NOT NULL REFERENCES hobi(id) ON DELETE CASCADE,           -- ID Hobi
    UNIQUE (id_mahasiswa, id_hobi)
    );

### Cara Menjalankan Aplikasi

go run main.go


### DOKUMENTASI API

POST	/mahasiswa	

Insert data mahasiswa baru.

GET	/mahasiswa	

Get semua data mahasiswa.

GET	/mahasiswa/:id	

Get detail data mahasiswa by ID.

POST	/jurusan	Insert data jurusan baru (tambahan).
POST	/hobi	Insert data hobi baru (tambahan)       SAYA BUTUH SYNTAX  .md nya 