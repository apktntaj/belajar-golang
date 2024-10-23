# WHAT I LEARN

## Duplicate Lines Program

```go
func DupLinesV2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("Tidak ada file")
	}

	for _, file := range files {
		c, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Message: %v\n", err)
		}
		countLines(c, counts)
		c.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
```

`fmt.Printf()` dan `fmt.Fprintf()` adalah dua fungsi dalam package fmt di Go yang digunakan untuk mencetak data dengan format, tetapi ada perbedaan utama dalam tujuan dan cara penggunaannya:

`fmt.Printf()` digunakan untuk menuliskan langsung ke standard output sedangkan `fmt.Fprintf()` menuliskan ke writer tertentu (misalnya file, buffer, atau koneksi jaringan).

Dalam sistem operasi, ada tiga aliran data standar yang digunakan oleh program untuk berkomunikasi dengan pengguna atau sistem lain:

1. stdin (standard input) – untuk menerima input.
2. stdout (standard output) – Untuk mengirimkan output normal selain error.
3. stderr (standard error) – Untuk mengirimkan pesan kesalahan.

Entah kenapa harus dipisah output normal dengan output error.

## io.Copy() & ioutil.ReadAll()

`io.Copy` akan langsung meneruskan `io.Reader` ke `io.Write` tanpa perlu menampung ke dalam array buffer. Sedangkan `ioutio.ReadAll()` menampung terlebih dahulu baru bisa diteruskan.

`io,Copy()` cocok untuk menampilkan file besar atau aliran data seperti unduhan file.

`ioutil.ReadAll()` lebih cocok untuk situasi di mana kamu perlu melakukan parsing JSON, memanipulasi teks, atau memodifikasi data sebelum menampilkannya

```go
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		b, err2 := ioutil.ReadAll(resp.Body)
		// _, err2 := io.Copy(os.Stdout, resp.Body)
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
```

## Bencmarking

Kita bisa menghitung lamanya sebuah function menjalankan komputasi dengan cara:

```go
start := time.Now()
duration := time.Since(start).Seconds()
```

## IO.DISCARD

Function `io.Copy(io.DIscard, resp.Body)` akan membaca data dari `resp.Body` dan membuangnya, tetapi tetap melaporkan jumlah byte yang berhasil dibaca.
