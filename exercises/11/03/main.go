package main

import "fmt"

type customErr struct {
	errType string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is an error with type: %v", ce.errType)
}

func main() {
	c1 := customErr{
		errType: "Custom",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
