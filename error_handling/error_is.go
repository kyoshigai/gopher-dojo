package main

import (
	"errors"
	"fmt"
)

type myError1 struct{}

func (me myError1) Error() string {
	return "my error 1"
}

func (me myError1) Is(target error) bool {
	return true
}

func main() {
	var err1 error = myError1{}
	var err2 error = errors.New("new error")
	fmt.Println(errors.Is(err1, err2))
}
