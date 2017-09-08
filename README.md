# binar-backend-assessment
Perancangan micro service untuk aplikasi pengantaran makanan minuman

Oleh: Yogi Agnia Dwi Saputro

## 1. Rancangan API
Desain API dibuat dengan layanan API blueprint Apiary. Alasannya adalah kemudahan penggunaan untuk keperluan data dummy serta dokumentasi yang langsung terstruktur.

Ada tiga entitas pokok terpenting dalam back end:
1. Pelanggan/customer  : ID, nama, email, nomor HP
2. Penjual/seller      : ID, nama, alamat
3. Pesanan/order       : ID, tanggal, pembayaran , pesanan(yang terdiri atas ID, nama, jumlah pesanan, harga)

Yang paling pokok adalah bagian order untuk mengetahui siapa yang memesan apa.

API blueprint dapat diakses di http://private-b0c8f3-binarfooddelivery.apiary-mock.com


## 2. Penanganan keamanan
Untuk memastikan keamanan data, proses koneksi harus diawali dengan autentikasi. Itu dimaksudkan agar penggunaan dapat dikontrol. Opsi untuk autentikasi ada:
1. Static API key
2. Simple authentication
3. OAuth authentication

Dalam contoh ini digunakan OAuth authentication

## 3. Service CRUD
Micro service dibuat menggunakan bahasa Go dengan library dasar seperti gin-gonic dan driver untuk database
Diasumsikan database telah tersedia.
