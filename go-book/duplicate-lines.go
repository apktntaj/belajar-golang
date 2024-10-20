package main

import (
	"bufio"
	"fmt"
	"os"
)

func DupLines() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

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
