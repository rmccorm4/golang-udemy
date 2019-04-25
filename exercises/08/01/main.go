package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	bs, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("test.json", bs, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
