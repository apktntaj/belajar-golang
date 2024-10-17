# Go Getting Started

```go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
```

- Selalu awali membuat proyek di Go dengan perintah `go mod init <nama_proyek>`
- `go run .` akan mengkompilasi kode sumber dan mengeksekusi program. Perilakunya mirip seperti bahasa yang di-interpretasi.
- `go help` adalah perintah untuk melihat command apa saja yang tersedia.
- Bila ingin menggunakan package dari pihak ketiga kita bisa menuliskan nama package pada saat deklarasi import lalu mengetikkan perintah `go mod tidy`.
