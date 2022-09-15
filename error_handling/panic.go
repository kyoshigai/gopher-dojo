package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func myPanic() {
	panic(errors.New("my panic"))
}

func f() (rerr error) {
	defer func() {
		if r := recover(); r != nil {
			if err, isErr := r.(error); isErr {
				rerr = err
			} else {
				rerr = fmt.Errorf("%v", r)
			}
		}
	}()
	// panic(errors.New("error in f"))
	myPanic()
	return nil
}

func main() {
	err := f()
	fmt.Println(err)
}
