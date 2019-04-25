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
	bs, err := ioutil.ReadFile("../01/test.json")
	var users []User
	err = json.Unmarshal(bs, &users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}
