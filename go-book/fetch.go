package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
