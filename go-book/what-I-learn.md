`fmt.Printf()` dan `fmt.Fprintf()` adalah dua fungsi dalam package fmt di Go yang digunakan untuk mencetak data dengan format, tetapi ada perbedaan utama dalam tujuan dan cara penggunaannya:

`fmt.Printf()` digunakan untuk menuliskan langsung ke standard output sedangkan `fmt.Fprintf()` menuliskan ke writer tertentu (misalnya file, buffer, atau koneksi jaringan).

Dalam sistem operasi, ada tiga aliran data standar yang digunakan oleh program untuk berkomunikasi dengan pengguna atau sistem lain:

1. stdin (standard input) – untuk menerima input.
2. stdout (standard output) – Untuk mengirimkan output normal selain error.
3. stderr (standard error) – Untuk mengirimkan pesan kesalahan.

Entah kenapa harus dipisah output normal dengan output error.
