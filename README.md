# Tugas-6-Pemrograman-Jaringan
### Javiar Fasyah         1301164477
### Hilmi Triandi N		    1301164286
### Fahrur Rozi Syarbini	1301164213  
  
  
**Jawaban Soal No. 1:**  
![fsm.png](/screenshot/fsm.png)  
Server membaca config dari file config.yaml untuk menentukan nama db, tabel, port db, dan port https untuk akses aplikasi. Setelah config dibaca, server mengecek koneksi dengan db yg digunakan apakah dapat dihubungi atau tidak, jika iya maka server akan membuka rute-rute yg dapat diakses pengguna dalam menggunakan aplikasi. Karena server melisten request https, maka dibutuhkan key dan certificate dengan menggunakan openssl, key dan certificate lalu dimasukan di dalam folder server. Server perlu dibuild dan dijalankan dalam mode sudo untuk bisa berjalan. Server pun siap listening dan aplikasi dapat diakses oleh pengguna.  
Pada aplikasi ini, pengguna dapat melakukan operasi CRUD seperti membuat post berita baru, melihat post berita, mengedit post berita, dan menghapus post berita. Diluar operasi CRUD, pengguna bisa memberi reaksi terhadap sebuah post berita dalam halaman melihat post berita. Operasi CRUD memanfaatkan rute-rute yg sudah dibuat server sebelumnya. Saat pengguna akan membuat post berita baru rutenya adalah “/new” lalu “/insert”, melihat post berita rutenya adalah “/view/{id berita}”, mengedit post berita rutenya adalah “/edit/{id berita}” lalu “/update”, dan menghapus berita adalah “/del/{id berita}”. Diluar rute-rute itu, pengguna akan diperingatkan bahwa halaman tersebut tidak ada.  
  
**Jawaban Soal No. 2:**  
Index:  
![index.png](/screenshot/index.png)  
  
New:  
![new.png](/screenshot/new.png)  
  
View:  
![view.png](/screenshot/view.png)  
  
Edit:  
![edit.png](/screenshot/edit.png)  
  
Delete:  
![delete.png](/screenshot/delete.png)  
  
*dalam tugas 7, server berjalan pada port 443 (tidak ditampilkan di browser) dan sertifikasinya sempat diragukan oleh browser (karena self-signed), bisa tetap lanjutkan ke dalam aplikasi web
