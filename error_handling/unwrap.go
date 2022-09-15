package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("bar: %w", errors.New("foo"))
	fmt.Println(err) // ber: foo
	err1 := errors.Unwrap(err)
	fmt.Println(err1) // foo
	err2 := errors.Unwrap(err1)
	fmt.Println(err2) // nil
	err3 := errors.Unwrap(err2)
	fmt.Println(err3) // nil
}
