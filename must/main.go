package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Must(checkNumber(10)))
	fmt.Println(Must(checkString("Nadhira")))
}
func Must(a any, err error) any {
	if err != nil {
		panic(err)
	}
	return a
}
func checkString(name string) (string, error) {
	if name != "Nadhira" {
		return name, errors.New("Not Nadhira")
	} else {
		return name, nil
	}

}

func checkNumber(a int) (int, error) {
	if a < 9 {
		return a, errors.New("A is less than 10")
	} else {
		return a, nil
	}
}
