# Revisi UAS-TubesPBW-Agen-Toko-DNK

## KELOMPOK 4 ##
## 1. Salwa Khairu Mista - 4522210066
## 2. Afni Puspita Zahra - 4522210120
## 3. Siti Ghumaisa      - 4522210131
   
---

# Deskripsi ##
TOKO DNK adalah aplikasi web yang dibuat untuk mendukung operasional toko agen yang menjual berbagai kebutuhan pokok (sembako). Melalui TOKO DNK, admin dapat dengan mudah mengelola produk, termasuk menambah, memperbarui, dan menghapus item dari inventaris toko dan pengguna dapat mengcheckout produk dari web ini. Diharapkan TOKO DNK dapat membantu meningkatkan efisiensi pengelolaan toko dan memastikan bahwa setiap proses berjalan dengan lancar dan terorganisir.

---

# Kerangka Pikiran
![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/38b40fd3-6a32-458f-ba69-61c60dde6192)
Alur singkatnya adalah pengguna memasukkan data (signup) dan masuk ke dalam database pengguna. Admin mengelola data produk (CRUD) di dalam database product. Pengguna dapat memasukkan barang, mengupdate kuantitas barang, dan melakukan transaksi pembelian. Dimana ketika memasukkan, mengupdate barang masuk ke dalam database cart. Kemudian, ketika pengguna melakukan transaksi maka akan masuk ke dalam database transactions.

---
# Aktor
1. Admin
2. User

# Database
1. admin
2. pengguna
3. product
4. cart
5. transactions

---
## Sebelumnya revisi di bagian LandingPage. Dikarenakan Landing Page antara admin dan user masih nyatu sebelumnya. Sekrang sudah dipisah.

---
# Tampilan Web
## 1. User
## ●	Sign Up User
 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/a43fb24a-2494-4ca7-a54e-05b159686520)
	Tampilan yang dimana user mendaftar username dan password yang akan nantinya disimpan di database.

---
## ●	Login User
 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/2543040d-da9f-415a-b62d-69b7beb056a2)
	Tampilan yang user pakai untuk masuk ke halaman LandingPage.

---
## ●	Landing Page User
 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/3e478dc7-e424-4694-9cf5-10c468bbd196)

 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/ed1a9a0c-cce0-44f6-b1fb-67c5b4f34dfb)
	Bagian ini adalah dimana adalah hasil create dari Admin ketika menambahkan produk.
 
 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/dffce818-c321-4d36-815c-308b4d9ac838)

 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/27b4bb25-b26a-4565-bdd9-998add6757a0)

 ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/68e6cc1a-05ea-4698-ac45-7ce0f66b59e5)
	Di LandingPage ini ada section Produk, Layanan Kami, dan Statistika Toko.

---
## ●	Tambah ke Keranjang/ create dari sisi user
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/11890d57-e996-484b-a067-c490206901f2)
	Pada LandingPage#unggulan ini menampilkan hasil dari create produk atau disebut view-nya. Dan dimana juga disini bisa menambahkan produk ke keranjang yang otomatis akan ke /cart.

---
## •	View Cart/read
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/3669c2bd-8caa-47bd-8b9e-12e3ca2fdf53)
	Halaman view cartnya, dimana ada kasi untuk meng-update dan delete kuantitas. Dan delete produk di keranjang.

---
## ●	Ubah Jumlah/Kuantitas Pembelian Produk/update
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/a31d6bdb-5e24-495a-8c0c-e048b408aad8)
	Halaman yang digunakan untuk meng-update jumlah barang di keranjang.

  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/aaa51fbd-0dff-42db-8842-67683f603226)
	Ketika sudah update jumlah kuantitasnya, maka otomatis juga harga total berubah.

---
## ●	Hapus produk yang tidak jadi dibeli/delete
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/78fabeeb-bd1f-416a-8ce6-9b72bfc80665)
	Dengan mengklik aksi hapus maka produk di keranjang akan otomatis terhapus.

  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/7bcfa1f2-f25c-40b7-8312-6114040adc7d)
	Tampilan ketika produk sudah dihapus.

---	
# Transaksi
## ●	Checkout 
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/0faee443-a5ff-47b8-9391-2948cd145eec)
	Klik ‘checkout’

  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/fbe25441-73bf-4364-9288-1ba059d95fc6)
	Maka keluar output yang sudah diinput dari keranjang beserta terdapat username di bagian ‘Terimakasih salwa telah berbelanja di toko kami’. Username diambil dari cookie ketika login.

---
# Tampilan Web
## Admin
## •	Signup Admin
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/785a258a-e796-4754-a883-d02a6efc965a)
	Tampilan yang dimana admin mendaftar username dan password yang akan nantinya disimpan di database.

 ---
## ●	Login Admin
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/6852ba8f-5b3d-435a-b953-403547c4edab)
	Tampilan yang admin pakai untuk masuk ke halaman LandingPage/admin.

---
## ●	Landing Page Admin
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/8e4bbd45-7eca-4839-8797-b9ae58d106f2)
	Bagian ini adalah dimana adalah hasil create dari Admin ketika menambahkan produk.

  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/ac5d63d7-9a43-48f5-b79d-1aa5e814e782)
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/2a6b27a8-80a0-43c4-89d3-34e5e5a36c00)
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/7d7de221-8f9b-4acd-8608-16c0e1543f84)
	Di LandingPage admin ini hanya ada, Layanan Kami, dan Statistika Toko yang memdekan dari Landing Page User serta pada landing admin terdapat button untuk mengarah ke create atau ke view product nya..

---
## ●	View List produk/read 
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/bdea35d8-494b-49ac-b13d-4b496ea6c8fa)
	Halaman ketika Admin sudah menambahkan produk, maka di halaman ini akan muncul produknya seperti view dari hasil tambah produk.

---
## ●	Jika ingin mengedit produk/update
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/3a0fa315-4756-4d27-8f83-1fc41c27755e)
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/47820f8f-d0d6-4d51-bec3-0b7886f2d146)
	Halaman untuk meng-update sebuah produk yang nantinya akan direct ke halaman /product/update

---
## ●	Jika ingin mendelete produk/delete
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/56fbe2cc-754e-40a0-9f4b-c92ecdfa4aa8)
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/8be016cb-4dea-4862-a705-82b048d6dab2)
  Delete produk dengan mengklik aksi delete yang dimana sudah terhubung ke databas, jadi ketika klik delete otomatis terhapus di database.

---
## ●	Tambahkan produk/create
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/c745bdca-2c80-4898-b40a-c4bcb6f2a298)
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/0b65f650-6669-4082-a422-ed5dc514bb0b)
	Halaman ketika Admin ingin menambahkan produk dan muncul di Listproduk serta muncul juga di user pada Landing Page . Berikut:
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/addccbba-7ae9-4cbb-8008-ebc07621cc05)

---
# Database
## Database dengan nama agen
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/e8e9295e-2461-4eba-8bd8-8f3f3b46ad37)
	Database yang mengelola tabel-tabel lain nantinya.

 ---
## •	Tabel Admin
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/3f33c32e-faca-41d2-86fa-f3048b8a7a6b)

---
## ●	Tabel pengguna
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/71c263af-25f4-4d63-bc37-1ea6184b83b4)
	Pada Tabel pengguna terdapat atribut id, username, dan password

---
## ●	Tabel product
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/3d759571-fbfe-4bb9-8ec5-be8ccc7c4f2c)
	Terdapat atribut id, nama, harga, deskripsi.

---
## ●	Tabel cart
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/d8592b39-35bc-47b6-a114-4c38604f6a45)
	Terdapat atribut id, product_name, kuantitas.

---
## ●	Tabel transactions
  ![image](https://github.com/afnipz/TubesPBW-Agen-Toko-DNK/assets/161412781/620788c8-82db-428c-8959-8a0b8995d2ef)
  Terdapat atribut id, product_name, quantity, unit_price, total_price, total_amount.

---


