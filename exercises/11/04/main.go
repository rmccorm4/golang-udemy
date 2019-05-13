package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-1)
	if err != nil {
		ce := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  err,
		}
		foo(ce)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		return -1, errors.New("Can't sqrt a negative number.")
	}
	return 42, nil
}

func foo(e error) {
	log.Println(e)
}
