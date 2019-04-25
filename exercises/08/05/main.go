package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (u byAge) Len() int           { return len(u) }
func (u byAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u byAge) Less(i, j int) bool { return u[i].Age < u[j].Age }

type byLast []user

func (u byLast) Len() int           { return len(u) }
func (u byLast) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u byLast) Less(i, j int) bool { return u[i].Last < u[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// your code goes here
	sort.Sort(byAge(users))

	for i, v := range users {
		sort.Strings(v.Sayings)
		fmt.Printf("%d\n\tFirst: %v\n\tLast: %v\n\tAge: %v\n\tSayings: %v\n", i, v.First, v.Last, v.Age, v.Sayings)
	}

	fmt.Println("----------")
	sort.Sort(byLast(users))

	for i, v := range users {
		sort.Strings(v.Sayings)
		fmt.Printf("%d\n\tFirst: %v\n\tLast: %v\n\tAge: %v\n\tSayings: %v\n", i, v.First, v.Last, v.Age, v.Sayings)
	}
}
