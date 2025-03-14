# Dokumentasi Tugas TEFA Golang Ke-3  

## Setup Proyek  

### 1. Inisialisasi Modul  
Jalankan perintah berikut untuk menginisialisasi proyek dan mengunduh dependensi:  
```sh
go mod init cashier
go get github.com/brianvoe/gofakeit/v6
go mod tidy
```

### 2. Menjalankan Aplikasi  
Gunakan perintah berikut untuk menjalankan aplikasi:  
```sh
go run main.go
```

## Contoh Output  
```
Kasir 1 melayani John Doe dengan 2 barang  
Kasir 2 melayani Jane Smith dengan 1 barang  
Kasir 3 melayani Mark Johnson dengan 3 barang  
Kasir 1 selesai melayani John Doe  
Kasir 4 melayani Emily Brown dengan 2 barang  
...
Semua transaksi selesai!
```
