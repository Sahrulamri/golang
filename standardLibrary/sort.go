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

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	user := []User{
		{Name: "Eko", Age: 10},
		{Name: "Kurniawan", Age: 200},
		{Name: "Khanedy", Age: 30},
	}

	sort.Sort(UserSlice(user))

	fmt.Println(user)
}