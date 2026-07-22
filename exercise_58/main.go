package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u User) Print() {
	fmt.Println(u.First, u.Last, u.Age)
	for _, s := range u.Sayings {
		fmt.Println("\t", s)
	}
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type ByAge struct {
	Users
}

func (u ByAge) Less(i, j int) bool {
	return u.Users[i].Age < u.Users[j].Age
}

type ByLast struct {
	Users
}

func (u ByLast) Less(i, j int) bool {
	return u.Users[i].Last < u.Users[j].Last
}

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}
	for _, u := range users {
		sort.Strings(u.Sayings)
		u.Print()
	}

	sort.Sort(ByAge{users})
	for _, u := range users {
		u.Print()
	}

	sort.Sort(ByLast{users})
	for _, u := range users {
		u.Print()
	}
}
