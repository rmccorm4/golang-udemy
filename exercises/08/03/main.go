package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "Ryan",
		Age:  22,
	}
	u2 := User{
		Name: "Jerry",
		Age:  26,
	}
	users := []User{u1, u2}
	ep := json.NewEncoder(os.Stdout)
	err := ep.Encode(users)
	if err != nil {
		panic(err)
	}
}
