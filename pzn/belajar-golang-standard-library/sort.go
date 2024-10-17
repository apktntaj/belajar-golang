package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// users := []User{
	// 	{"Eko", 30},
	// 	{"Budi", 35},
	// 	{"Joko", 25},
	// 	{"Adit", 20},
	// }

	var users UserSlice
	users = UserSlice{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 20},
	}

	// sort.Sort(UserSlice(users))
	sort.Sort(users)

	fmt.Println(users)
}
