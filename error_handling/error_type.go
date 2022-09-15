package main

import (
	"log"
	"os"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func doSomething2() error {
	path := "./hoge.txt"
	f, err := os.Open(path)
	if err != nil {
		return &PathError{
			Op:   "doSomething2",
			Path: path,
			Err:  err,
		}
	}
	defer f.Close()

	// do something

	return nil
}

func main() {
	if err := doSomething2(); err != nil {
		log.Println(err)
	} else {
		log.Println("ok")
	}
}
