package main

import "efishery-assignment/infrastructure/router"

// Temen-temen sudah dapat tema ya, temanya e-commerce. Tugasnya bikin REST API yang isinya:

// 1. Detail produk -> nama , foto, harga, deskripsi
// 2. Filter produk -> berdasarkan kategori, berdasarkan range harga
// 3. Lissting produk -> nama, foto, harga, kategori
// 4. Pembayaran -> bayar sesuai total pembelian dan upload bukti pembayaran (cukup upload saja, tidak perlu verifikasi)
// 5. Tambahkan produk ke troli atau cart

// Catatan:
// 1. Dokumentasi REST API bisa pakai swagger atau buat tabel di spreadsheet
// 2. Tidak perlu UI. Semu via REST API
// 3. Demo bisa menggunakan Postman
// 4. Lama pengerjaan 1 minggu
// 5. Akan ada sesi speeddating dengan eFisherian dan User

func main() {
	router.Init()
}
