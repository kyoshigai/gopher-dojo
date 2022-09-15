package main

import (
	"errors"
	"log"
	"os"
)

func doSomething(f *os.File) error {
	_ = f
	return errors.New("doSomething error")
}

func main() {
	var err error
	f, err := os.Open("not_exist.txt")
	if err != nil {
		log.Println("1番目のエラーハンドリング", err)
	}
	defer f.Close()

	// 本来は err = doSomething(f) としたつもり
	doSomething(f)
	if err != nil {
		log.Println("2番目のエラーハンドリング", err)
	}
}
