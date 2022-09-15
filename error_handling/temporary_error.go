package main

import "fmt"

type temporary interface {
	Temporary() bool
}

type tmp struct{}

func (t tmp) Temporary() bool {
	return true
}
func (t tmp) Error() string {
	return "temp error"
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

func main() {
	t := &tmp{}
	fmt.Println(IsTemporary(t))
}
