package main

import (
	"fmt"
	"os"
	"strings"
)

func Arguments() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func ArgumentsV2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Modify the echo program to also print os.Args[0]
func ExcerciseArgs1() {
	for _, arg := range os.Args {

		fmt.Println(arg)
	}
}

// Modify the echo program to print the index and value of each of its arguments
func ExcerciseArgs2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d. %s\n", i+1, arg)
	}
}

// Using strings.Join
func ExcerciseArgs3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
