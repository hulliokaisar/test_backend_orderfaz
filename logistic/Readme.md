# Logistic App

Logistic App adalah sebuah aplikasi yang digunakan untuk mengelola data kurir. Aplikasi ini dibangun menggunakan Golang dan menggunakan database SQL untuk menyimpan data kurir.

## Fitur

- Mengambil data kurir dari database
- Melakukan filter data kurir berdasarkan origin_name dan destination_name

## Instalasi

1. Pastikan Anda telah menginstal Go (Golang) dan Docker di komputer Anda.
2. Clone repositori ini ke direktori lokal Anda.
3. Aplikasi akan diakses melalui `localhost:8000`.

## Penggunaan

- Untuk mengambil data semua kurir, jalankan `GET` request ke `http://localhost:8000/couriers`.
- Untuk melakukan filter data kurir berdasarkan origin_name dan destination_name, jalankan `GET` request ke `http://localhost:8000/couriers/{origin_name}/{destination_name}`.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan melakukan fork repositori ini, buat branch untuk perubahan Anda, dan buat pull request ke branch `main` dari repositori ini.

## Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT. Lihat berkas `LICENSE` untuk informasi lebih lanjut.
